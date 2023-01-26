package db

import (
	"errataService/pkg/configurator"
	"errataService/pkg/utils"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2"
	"net/http"
	"strconv"
	"time"
)

func InitDB() (*DB, error) {
	opt := clickhouse.Options{
		Addr: []string{configurator.Config.AddressToClick},
		Auth: clickhouse.Auth{
			Database: configurator.Config.DataBase,
			Username: configurator.Config.Login,
			Password: configurator.Config.Password,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: time.Duration(configurator.Config.DialTimeout) * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	}
	if configurator.Config.HTTP {
		opt.Protocol = clickhouse.HTTP
	}
	conn := clickhouse.OpenDB(&opt)
	ping := conn.Ping()
	if ping != nil {
		return nil, ping
	}
	dataBase := DB{
		db: conn,
	}
	return &dataBase, nil
}

func (db *DB) CheckConnect() bool {
	status := db.db.Ping()
	if status == nil {
		return true
	}
	return false
}

func (db *DB) GetErrata(errata_id string) (*Errata, int, error) {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	var errata Errata
	row := db.db.QueryRow("SELECT * FROM errata WHERE errata_id = $1", errata_id)
	if err := row.Scan(&errata.id, &errata.Prefix, &errata.Num, &errata.UpdateCount, &errata.CreationDate, &errata.ChangeDate); err != nil {
		return nil, http.StatusNotFound, err
	}
	return &errata, http.StatusOK, nil
}

func (db *DB) UpdateErrata(errata_id string, update int64) (*Errata, int, error) {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	var errata Errata
	row := db.db.QueryRow("SELECT * FROM errata WHERE errata_id = $1 AND errata_update_count = (SELECT max(errata_update_count) FROM errata WHERE errata_id= $1)", errata_id)
	if err := row.Scan(&errata.id, &errata.Prefix, &errata.Num, &errata.UpdateCount, &errata.CreationDate, &errata.ChangeDate); err != nil {
		return nil, http.StatusNotFound, err
	}
	if errata.UpdateCount != update {
		return nil, http.StatusNotFound, errors.New("version don`t match")
	}
	errata.UpdateCount += 1
	errata.ChangeDate = time.Now()
	_, err := db.db.Exec("INSERT INTO errata VALUES ($1, $2, $3, $4, $5, $6)",
		errata.id, errata.Prefix, errata.Num, errata.UpdateCount,
		errata.CreationDate, errata.ChangeDate)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return &errata, http.StatusOK, nil
}

func (db *DB) GenerateErrata(prefix string) (*Errata, int, error) {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	var last int64
	var current int64
	row := db.db.QueryRow("SELECT max(errata_num) FROM errata")
	if err := row.Scan(&last); err != nil {
		return nil, http.StatusNotFound, err
	}
	if last < 999 {
		last = 999
	}
	current = last + 1
	id := utils.SHA1(prefix + "-" + strconv.Itoa(int(current)))
	errata := CreateErrata(id, prefix, current, 1, time.Now(), time.Now())
	_, err := db.db.Exec("INSERT INTO errata VALUES ($1, $2, $3, $4, $5, $6)",
		errata.id, errata.Prefix, errata.Num, errata.UpdateCount,
		errata.CreationDate, errata.ChangeDate)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return errata, http.StatusOK, nil
}

func (db *DB) Close() {
	db.db.Close()
}

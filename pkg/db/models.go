package db

import (
	"database/sql"
	"sync"
	"time"
)

type DB struct {
	db  *sql.DB
	mtx sync.Mutex
}

type Errata struct {
	id           string
	Prefix       string    `json:"prefix"`
	Num          int64     `json:"num"`
	UpdateCount  int64     `json:"updateCount"`
	CreationDate time.Time `json:"creationDate"`
	ChangeDate   time.Time `json:"changeDate"`
}

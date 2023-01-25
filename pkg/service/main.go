package service

import (
	"errataService/pkg/configurator"
	"errataService/pkg/db"
	"errataService/pkg/logger"
	"time"
)

var Service ServiceT

var infoLogger = logger.NewLogger("INFO")
var errorLogger = logger.NewLogger("ERROR")

func CreateService() error {
	conn, err := db.InitDB()
	for err != nil {
		errorLogger.Printf("connection to the database failed: %s", err.Error())
		conn, err = db.InitDB()
		time.Sleep(time.Second * 4)
	}
	infoLogger.Printf("success connect to Clickhouse: %s", configurator.Config.AddressToClick)
	Service = ServiceT{
		db: conn,
	}
	return nil
}

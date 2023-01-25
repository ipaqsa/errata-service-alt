package server

import (
	"errataService/pkg/configurator"
	"errataService/pkg/logger"
	"errataService/pkg/service"
	"net/http"
)

var infoLogger = logger.NewLogger("INFO")
var errorLogger = logger.NewLogger("ERROR")

func Run() error {
	err := service.CreateService()
	if err != nil {
		return err
	}
	defer service.Service.CloseConnect()

	http.HandleFunc("/errata", errataHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/check", checkHandler)

	infoLogger.Printf("Service start at %s", configurator.Config.Port)
	err = http.ListenAndServe(configurator.Config.Port, nil)
	if err != nil {
		return err
	}
	return nil
}

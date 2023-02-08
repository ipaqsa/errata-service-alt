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

	http.HandleFunc("/register", errataHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/check", checkHandler)

	infoLogger.Printf("Service start at %s", configurator.Port)
	err = http.ListenAndServe(configurator.Port, nil)
	if err != nil {
		return err
	}
	return nil
}

package server

import (
	"errataService/pkg/logger"
	"errataService/pkg/service"
	"net/http"
)

var infoLogger = logger.NewLogger("INFO")
var errorLogger = logger.NewLogger("ERROR")

func Run(port string) error {
	err := service.CreateService()
	if err != nil {
		return err
	}
	defer service.Service.CloseConnect()

	http.HandleFunc("/register", errataHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/check", checkHandler)

	infoLogger.Printf("Service start at %s", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}
	return nil
}

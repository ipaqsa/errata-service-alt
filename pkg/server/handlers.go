package server

import (
	"errataService/pkg/configurator"
	"errataService/pkg/service"
	"errataService/pkg/utils"
	"net/http"
	"strings"
)

func errataHandler(w http.ResponseWriter, r *http.Request) {
	splits := strings.Split(r.RemoteAddr, ":")
	if len(splits) != 2 || !utils.Contains(configurator.Config.Allowed, splits[0]) {
		errorLogger.Printf("Dont allowed host: %s", splits[0])
		for _, addr := range configurator.Config.Allowed {
			println(addr)
		}
		err := sendAnswer(w, -1, "Access denied", nil)
		if err != nil {
			errorLogger.Printf(err.Error())
		}
		return
	}
	if r.Method == http.MethodGet {
		request, err := UnmarshalRequest(r)
		if err != nil {
			errorLogger.Printf(err.Error())
			err := sendAnswer(w, -1, "wrong request", nil)
			if err != nil {
				errorLogger.Printf(err.Error())
			}
			return
		}
		errata, err := service.Service.GenerateErrata(request.Data)
		if err != nil {
			errorLogger.Printf(err.Error())
			err := sendAnswer(w, -1, err.Error(), nil)
			if err != nil {
				errorLogger.Printf(err.Error())
			}
			return
		}
		err = sendAnswer(w, 1, "OK", errata)
		if err != nil {
			errorLogger.Printf(err.Error())
			return
		}
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	splits := strings.Split(r.RemoteAddr, ":")
	if len(splits) != 2 || !utils.Contains(configurator.Config.Allowed, splits[0]) {
		errorLogger.Printf("Dont allowed host: %s", splits[0])
		err := sendAnswer(w, -1, "Access denied", nil)
		if err != nil {
			errorLogger.Printf(err.Error())
		}
		return
	}
	if r.Method == http.MethodPost {
		request, err := UnmarshalRequest(r)
		if err != nil {
			errorLogger.Printf(err.Error())
			err := sendAnswer(w, -1, "wrong request", nil)
			if err != nil {
				errorLogger.Printf(err.Error())
			}
			return
		}
		errata, err := service.Service.UpdateErrata(request.Data)
		if err != nil {
			errorLogger.Printf(err.Error())
			err := sendAnswer(w, -1, err.Error(), nil)
			if err != nil {
				errorLogger.Printf(err.Error())
			}
			return
		}
		err = sendAnswer(w, 1, "OK", errata)
		if err != nil {
			errorLogger.Printf(err.Error())
			return
		}
	}
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	splits := strings.Split(r.RemoteAddr, ":")
	if len(splits) != 2 || !utils.Contains(configurator.Config.Allowed, splits[0]) {
		errorLogger.Printf("Dont allowed host: %s", splits[0])
		err := sendAnswer(w, -1, "Access denied", nil)
		if err != nil {
			errorLogger.Printf(err.Error())
		}
		return
	}
	if r.Method == http.MethodGet {
		request, err := UnmarshalRequest(r)
		if err != nil {
			errorLogger.Printf(err.Error())
			err := sendAnswer(w, -1, "wrong request", nil)
			if err != nil {
				errorLogger.Printf(err.Error())
			}
			return
		}
		errata, err := service.Service.GetErrata(request.Data)
		if err != nil {
			errorLogger.Printf(err.Error())
			if err.Error() == "sql: no rows in result set" {
				err = sendAnswer(w, -1, "Don`t find", nil)
			} else {
				err = sendAnswer(w, -1, err.Error(), nil)
			}
			if err != nil {
				errorLogger.Printf(err.Error())
			}
			return
		}
		err = sendAnswer(w, 1, "OK", errata)
		if err != nil {
			errorLogger.Printf(err.Error())
			return
		}
	}
}

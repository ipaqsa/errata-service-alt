package server

import (
	"errataService/pkg/service"
	"net/http"
)

func errataHandler(w http.ResponseWriter, r *http.Request) {
	if !accessAddress(w, r) {
		return
	}
	if r.Method == http.MethodGet {
		prefix, s, err := parseQuery(r)
		if err != nil {
			errorLogger.Printf(err.Error())
			err = sendAnswer(w, s, "wrong request", nil)
			if err != nil {
				errorLogger.Printf(err.Error())
			}
			return
		}
		errata, s, err := service.Service.GenerateErrata(prefix)
		if err != nil {
			errorLogger.Printf(err.Error())
			err = sendAnswer(w, s, err.Error(), nil)
			if err != nil {
				errorLogger.Printf(err.Error())
			}
			return
		}
		err = sendAnswer(w, s, "OK", errata)
		if err != nil {
			errorLogger.Printf(err.Error())
			return
		}
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if !accessAddress(w, r) {
		return
	}
	if r.Method == http.MethodPost {
		name, s, err := parseQuery(r)
		if err != nil {
			errorLogger.Printf(err.Error())
			err = sendAnswer(w, s, "wrong request", nil)
			if err != nil {
				errorLogger.Printf(err.Error())
			}
			return
		}
		errata, s, err := service.Service.UpdateErrata(name)
		if err != nil {
			errorLogger.Printf(err.Error())
			if err.Error() == "sql: no rows in result set" {
				err = sendAnswer(w, http.StatusNotFound, "not found", nil)
			} else {
				err = sendAnswer(w, s, err.Error(), nil)
			}
			if err != nil {
				errorLogger.Printf(err.Error())
			}
			return
		}
		err = sendAnswer(w, s, "OK", errata)
		if err != nil {
			errorLogger.Printf(err.Error())
			return
		}
	}
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	if !accessAddress(w, r) {
		return
	}
	if r.Method == http.MethodGet {
		name, s, err := parseQuery(r)
		if err != nil {
			errorLogger.Printf(err.Error())
			err = sendAnswer(w, s, "wrong request", nil)
			if err != nil {
				errorLogger.Printf(err.Error())
			}
			return
		}
		errata, s, err := service.Service.GetErrata(name)
		if err != nil {
			errorLogger.Printf(err.Error())
			if err.Error() == "sql: no rows in result set" {
				err = sendAnswer(w, http.StatusNotFound, "not found", nil)
			} else {
				err = sendAnswer(w, s, err.Error(), nil)
			}
			if err != nil {
				errorLogger.Printf(err.Error())
			}
			return
		}
		err = sendAnswer(w, s, "OK", errata)
		if err != nil {
			errorLogger.Printf(err.Error())
			return
		}
	}
}

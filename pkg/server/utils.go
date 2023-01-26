package server

import (
	"encoding/json"
	"errataService/pkg/configurator"
	"errataService/pkg/db"
	"errataService/pkg/utils"
	"errors"
	"net/http"
	"regexp"
	"strings"
)

func valid(data string, vtype string) (string, bool) {
	if vtype == "prefix" {
		matched, _ := regexp.MatchString("^[A-Z]+[\\-0-9A-Z]+$", data)
		return data, matched
	} else if vtype == "name" {
		matched, _ := regexp.MatchString("^[A-Z]+[\\-0-9A-Z]+-[\\d]{4,}-[\\d]{1,}$", data)
		return data, matched
	}
	return "", false
}

func parseQuery(r *http.Request) (string, int, error) {
	q := r.URL.RawQuery
	splits := strings.Split(q, "=")
	if len(splits) != 2 {
		return "", http.StatusBadRequest, errors.New("wrong format")
	}
	data, status := valid(splits[1], splits[0])
	if !status {
		return "", http.StatusBadRequest, errors.New("wrong format")
	}
	return data, http.StatusOK, nil
}

func marshalResponse(response *Response) ([]byte, error) {
	marshal, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func sendAnswer(w http.ResponseWriter, status int, Comment string, errata *db.Errata) error {
	var respErrata ResponseErrata
	if errata != nil {
		respErrata = ResponseErrata{
			Errata:  db.ErrataToString(errata),
			Created: errata.CreationDate,
			Updated: errata.ChangeDate,
		}
	}
	resp := Response{
		Comment: Comment,
		Errata:  &respErrata,
	}
	response, err := marshalResponse(&resp)
	if err != nil {
		return err
	}
	w.WriteHeader(status)
	_, err = w.Write(response)
	if err != nil {
		return err
	}
	return nil
}

func accessAddress(w http.ResponseWriter, r *http.Request) bool {
	addr := getAddress(r)
	splits := strings.Split(addr, ":")
	if len(splits) != 2 || !utils.Contains(configurator.Config.Allowed, splits[0]) {
		errorLogger.Printf("Dont allowed host: %s", splits[0])
		err := sendAnswer(w, http.StatusForbidden, "Access denied", nil)
		if err != nil {
			errorLogger.Printf(err.Error())
		}
		return false
	}
	return true
}

func getAddress(r *http.Request) string {
	var userIP string
	if len(r.Header.Get("CF-Connecting-IP")) > 1 {
		userIP = r.Header.Get("CF-Connecting-IP")
		return userIP
	} else if len(r.Header.Get("X-Forwarded-For")) > 1 {
		userIP = r.Header.Get("X-Forwarded-For")
		return userIP
	} else if len(r.Header.Get("X-Real-IP")) > 1 {
		userIP = r.Header.Get("X-Real-IP")
		return userIP
	} else {
		return r.RemoteAddr
	}
}

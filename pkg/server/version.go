package server

import (
	"encoding/json"
	"errataService/pkg/configurator"
	"net/http"
)

type versionResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func versionSendAnswer(w http.ResponseWriter, status int) error {
	var respVersion versionResponse
	if status == http.StatusOK {
		respVersion = versionResponse{
			Name:    configurator.GetName(),
			Version: configurator.GetVersion(),
		}
	} else {
		respVersion = versionResponse{
			Name:    "",
			Version: "",
		}
	}
	response, err := json.Marshal(respVersion)
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

func versionHandler(w http.ResponseWriter, r *http.Request) {
	if !accessAddress(w, r) {
		return
	}
	if r.Method == http.MethodGet {
		err := versionSendAnswer(w, http.StatusOK)
		if err != nil {
			errorLogger.Printf(err.Error())
			return
		}
	} else {
		err := versionSendAnswer(w, http.StatusMethodNotAllowed)
		if err != nil {
			errorLogger.Println(err.Error())
			return
		}
	}
}

package server

import (
	"encoding/json"
	"errataService/pkg/db"
	"net/http"
)

func UnmarshalRequest(r *http.Request) (*RequestErrata, error) {
	decoder := json.NewDecoder(r.Body)
	var req RequestErrata
	err := decoder.Decode(&req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func MarshalResponse(response *ResponseErrata) ([]byte, error) {
	marshal, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func sendAnswer(w http.ResponseWriter, status int, statusData string, data *db.Errata) error {
	resp := ResponseErrata{
		Status:     status,
		StatusData: statusData,
		Errata:     data,
	}
	response, err := MarshalResponse(&resp)
	if err != nil {
		return err
	}
	_, err = w.Write(response)
	if err != nil {
		return err
	}
	return nil
}

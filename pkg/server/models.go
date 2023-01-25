package server

import "errataService/pkg/db"

type RequestErrata struct {
	Data string `json:"Data"`
}

type ResponseErrata struct {
	Status     int        `json:"status"`
	StatusData string     `json:"statusData"`
	Errata     *db.Errata `json:"errata"`
}

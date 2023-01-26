package server

import (
	"time"
)

type Response struct {
	Comment string          `json:"comment"`
	Errata  *ResponseErrata `json:"errata"`
}

type ResponseErrata struct {
	Errata  string    `json:"id"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

package server

type Response struct {
	Comment string          `json:"comment"`
	Errata  *ResponseErrata `json:"errata"`
}

type ResponseErrata struct {
	Errata  string `json:"id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

package db

import (
	"database/sql"
	"sync"
	"time"
)

type DB struct {
	db  *sql.DB
	mtx sync.Mutex
}

type Errata struct {
	id           string
	Prefix       string    `json:"prefix"`
	Year         uint32    `json:"year"`
	Num          uint32    `json:"num"`
	UpdateCount  uint32    `json:"updateCount"`
	CreationDate time.Time `json:"creationDate"`
	ChangeDate   time.Time `json:"changeDate"`
}

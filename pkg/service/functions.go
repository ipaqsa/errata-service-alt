package service

import (
	"errataService/pkg/db"
	"errors"
	"time"
)

func (service *ServiceT) GenerateErrata(prefix string) (*db.Errata, error) {
	status := service.db.CheckConnect()
	if !status {
		time.Sleep(time.Second)
		err := service.tryConnect()
		if err != nil {
			return nil, err
		}
		status = service.db.CheckConnect()
		if !status {
			return nil, errors.New("connection to the database failed")
		}
	}
	return service.db.GenerateErrata(prefix)
}

func (service *ServiceT) GetErrata(errata string) (*db.Errata, error) {
	status := service.db.CheckConnect()
	if !status {
		time.Sleep(time.Second)
		err := service.tryConnect()
		if err != nil {
			return nil, err
		}
		status = service.db.CheckConnect()
		if !status {
			return nil, errors.New("connection to the database failed")
		}
	}

	errata_id, err := db.ErrataToID(errata)
	if err != nil {
		return nil, err
	}
	return service.db.GetErrata(errata_id)
}

func (service *ServiceT) UpdateErrata(errata string) (*db.Errata, error) {
	status := service.db.CheckConnect()
	if !status {
		time.Sleep(time.Second)
		err := service.tryConnect()
		if err != nil {
			return nil, err
		}
		status = service.db.CheckConnect()
		if !status {
			return nil, errors.New("connection to the database failed")
		}
	}
	errata_id, err := db.ErrataToID(errata)
	if err != nil {
		return nil, err
	}
	return service.db.UpdateErrata(errata_id)
}

func (service *ServiceT) CloseConnect() {
	service.db.Close()
}

func (service *ServiceT) tryConnect() error {
	conn, err := db.InitDB()
	if err != nil {
		return err
	}
	service.db = conn
	return nil
}

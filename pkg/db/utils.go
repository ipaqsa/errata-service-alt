package db

import (
	"errataService/pkg/utils"
	"errors"
	"fmt"
	"strings"
	"time"
)

func CreateErrata(id, prefix string, num, updateCount int64, creationDate, changeDate time.Time) *Errata {
	return &Errata{
		id:           id,
		Prefix:       prefix,
		Num:          num,
		UpdateCount:  updateCount,
		CreationDate: creationDate,
		ChangeDate:   changeDate,
	}
}

func PrintErrata(errata *Errata) {
	fmt.Printf("Errata #%s %s-%d-%d Created: %s Last Change: %s\n", errata.id, errata.Prefix, errata.Num, errata.UpdateCount, errata.CreationDate, errata.ChangeDate)
}

func ErrataToID(errata string) (string, error) {
	splits := strings.Split(errata, "-")
	if len(splits) != 2 {
		return "", errors.New("wrong format, need: PREFIX-NUM")
	}
	return utils.SHA1(strings.Join(splits[:2], "-")), nil
}

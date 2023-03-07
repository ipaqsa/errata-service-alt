package db

import (
	"errataService/pkg/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CreateErrata(id, prefix, year string, num, updateCount int64, creationDate, changeDate time.Time) *Errata {
	return &Errata{
		id:           id,
		Prefix:       prefix,
		Year:         year,
		Num:          num,
		UpdateCount:  updateCount,
		CreationDate: creationDate,
		ChangeDate:   changeDate,
	}
}

func ErrataToString(errata *Errata) string {
	return errata.Prefix + "-" + errata.Year + "-" + strconv.FormatInt(errata.Num, 10) + "-" + strconv.FormatInt(errata.UpdateCount, 10)
}

func PrintErrata(errata *Errata) {
	fmt.Printf("Errata #%s %s-%s-%d-%d Created: %s Last Change: %s\n", errata.id, errata.Prefix, errata.Year, errata.Num, errata.UpdateCount, errata.CreationDate, errata.ChangeDate)
}

func ErrataToID(errata string) (string, int64, error) {
	splits := strings.Split(errata, "-")
	if len(splits) < 3 {
		return "", 0, errors.New("wrong format, need: PREFIX-NUM-YEAR-UPDATE")
	}
	update, err := strconv.ParseInt(splits[len(splits)-1], 10, 64)
	if err != nil {
		return "", 0, errors.New("wrong format, need: PREFIX-NUM-YEAR-UPDATE")
	}
	if update == 0 {
		return "", 0, errors.New("updated count cannot be equal to 0")
	}
	return utils.SHA1(strings.Join(splits[:len(splits)-1], "-")), update, nil
}

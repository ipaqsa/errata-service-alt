package server

import (
	"encoding/json"
	"errataService/pkg/configurator"
	"errataService/pkg/db"
	"errataService/pkg/utils"
	"errors"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func valid(data string, vtype string) (string, bool) {
	if vtype == "prefix" {
		matched, _ := regexp.MatchString("^[A-Z]+[\\-0-9A-Z]+$", data)
		return data, matched
	} else if vtype == "name" {
		matched, _ := regexp.MatchString("^[A-Z]+[\\-0-9A-Z]+-[\\d]{4,}-[\\d]{1,}$", data)
		return data, matched
	} else if vtype == "year" {
		matched, _ := regexp.MatchString("^[2][0-9]{3}$", data)
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

func parseRegisterQuery(r *http.Request) (*RegisterQuery, int, error) {
	q := r.URL.RawQuery
	splits := strings.Split(q, "&")
	if len(splits) != 2 {
		errorLogger.Printf("request invalid")
		return nil, http.StatusBadRequest, errors.New("wrong format")
	}
	var rq RegisterQuery
	prefixSplits := strings.Split(splits[0], "=")
	if len(prefixSplits) != 2 {
		errorLogger.Printf("request invalid")
		return nil, http.StatusBadRequest, errors.New("wrong format")
	}
	prefix, status := valid(prefixSplits[1], prefixSplits[0])
	if !status {
		errorLogger.Printf("%s invalid", prefixSplits[0])
		return nil, http.StatusBadRequest, errors.New("wrong format")
	}
	if prefixSplits[0] == "prefix" {
		rq.Prefix = prefix
	} else {
		yearUint, err := strconv.ParseUint(prefix, 10, 32)
		if err != nil {
			errorLogger.Printf("year parse error")
			return nil, http.StatusBadRequest, errors.New("wrong format")
		}
		rq.Year = uint32(yearUint)
	}
	yearSplits := strings.Split(splits[1], "=")
	year, status := valid(yearSplits[1], yearSplits[0])
	if !status {
		errorLogger.Printf("%s invalid", yearSplits[0])
		return nil, http.StatusBadRequest, errors.New("wrong format")
	}
	if yearSplits[0] == "year" {
		yearUint, err := strconv.ParseUint(year, 10, 32)
		if err != nil {
			errorLogger.Printf("year parse error")
			return nil, http.StatusBadRequest, errors.New("wrong format")
		}
		rq.Year = uint32(yearUint)
	} else {
		rq.Prefix = year
	}

	return &rq, http.StatusOK, nil
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
			Created: errata.CreationDate.Format("2006-01-02T15:04:05.000-07:00"),
			Updated: errata.ChangeDate.Format("2006-01-02T15:04:05.000-07:00"),
		}
	}
	if errata == nil {
		respErrata = ResponseErrata{
			Errata:  "",
			Created: "",
			Updated: "",
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
	nets := getNetworks(configurator.Config.Allowed)
	if len(splits) != 2 || !(utils.Contains(configurator.Config.Allowed, splits[0]) || inNetworks(nets, splits[0])) {
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

func getNetworks(allowed []string) []*net.IPNet {
	var data []*net.IPNet
	for _, allow := range allowed {
		_, ntw, err := net.ParseCIDR(allow)
		if err != nil {
			continue
		}
		data = append(data, ntw)
	}
	return data
}

func inNetworks(networks []*net.IPNet, addr string) bool {
	for _, ntw := range networks {
		if ntw.Contains(net.ParseIP(addr)) {
			return true
		}
	}
	return false
}

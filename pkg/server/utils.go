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
	data = strings.TrimSpace(data)
	if data == "" {
		return "", false
	}
	if vtype == "prefix" {
		matched, _ := regexp.MatchString("[A-Z]+-[A-Z]+$", data)
		return data, matched
	} else if vtype == "name" {
		matched, _ := regexp.MatchString("^[A-Z]+-[A-Z]+-2[\\d]{3}-[\\d]{4,}-[\\d]{1,4}$", data)
		return data, matched
	} else if vtype == "year" {
		matched, _ := regexp.MatchString("^2[0-9]{3}$", data)
		return data, matched
	}
	return "", false
}

func parseQuery(r *http.Request) (string, int, error) {
	q := r.URL.Query()
	name, status := valid(q.Get("name"), "name")
	if !status {
		return "", http.StatusBadRequest, errors.New("wrong name format")
	}
	return name, http.StatusOK, nil
}

func parseRegisterQuery(r *http.Request) (string, uint32, int, error) {
	qp := r.URL.Query()
	prefix, status := valid(qp.Get("prefix"), "prefix")
	if !status {
		return "", 0, http.StatusBadRequest, errors.New("wrong prefix format")
	}
	year, status := valid(qp.Get("year"), "year")
	if !status {
		return "", 0, http.StatusBadRequest, errors.New("wrong year format")
	}
	parseUint, err := strconv.ParseUint(year, 10, 32)
	if err != nil {
		return "", 0, http.StatusBadRequest, errors.New("wrong year format")
	}
	return prefix, uint32(parseUint), http.StatusOK, nil
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

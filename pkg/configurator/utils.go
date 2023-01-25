package configurator

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net"
	"strings"
)

func pathToConfIsValid(path string) error {
	splits := strings.Split(path, ".")
	if len(splits) < 2 {
		return errors.New("path invalid")
	}
	format := splits[len(splits)-1]
	if format != "yml" {
		return errors.New("format is not yml")
	}
	return nil
}

func parseConfig(path string) (*ConfigT, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config ConfigT
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	if config.DialTimeout < 1 {
		config.DialTimeout = 5
	}
	return &config, nil
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

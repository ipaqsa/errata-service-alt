package configurator

import (
	"errors"
	"io/ioutil"
	"net"
	"strings"

	"gopkg.in/yaml.v3"
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
	if config.Name == "" {
		config.Name = defaultName
	}
	if config.Port == 0 {
		config.Port = defaultPort
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

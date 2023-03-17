package main

import (
	"errataService/pkg/configurator"
	"errataService/pkg/server"
	"os"
)

var version = ""

func init() {
	err := configurator.FlagInit()
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}
	configurator.SetVersion(version)
	configurator.InitInfo()
	configurator.PrintInfo()
}

func main() {
	err := server.Run()
	if err != nil {
		println(err.Error())
		return
	}
}

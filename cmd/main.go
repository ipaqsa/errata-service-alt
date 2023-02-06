package main

import (
	"errataService/pkg/configurator"
	"errataService/pkg/server"
	"os"
)

func init() {
	err := configurator.FlagInit()
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}
	configurator.SetVersion("0.0.1")
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

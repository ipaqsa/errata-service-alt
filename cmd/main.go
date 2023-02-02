package main

import (
	"errataService/pkg/configurator"
	"errataService/pkg/server"
	"os"
)

const port = "9111"

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
	println("Hello!")
	err := server.Run(port)
	if err != nil {
		println(err.Error())
		return
	}
}

package configurator

import (
	"errors"
	"flag"
)

var PathToConfig *string
var Debug *bool

var Config *ConfigT

func FlagInit() error {
	PathToConfig = flag.String("c", "", "Path to config file")
	Debug = flag.Bool("b", false, "Debug mode")
	flag.Parse()
	if *PathToConfig != "" {
		err := pathToConfIsValid(*PathToConfig)
		if err != nil {
			return err
		}
		Config, err = parseConfig(*PathToConfig)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("add flag `-c` - path to config")
}

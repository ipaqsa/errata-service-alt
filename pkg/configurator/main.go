package configurator

import (
	"fmt"
	"os"
	"runtime"
)

var Info InfoT

func InitInfo() {
	Info.Name = Config.Name
	Info.OS = runtime.GOOS
	Info.Arch = runtime.GOARCH
	if _, err := os.Lstat("/.dockerenv"); err != nil && os.IsNotExist(err) {
		Info.Container = "outside"
	} else {
		Info.Container = "inside"
	}
	Info.Hostname, _ = os.Hostname()
	taddr := getOutboundIP()
	if taddr == nil {
		Info.Address = "internal network is not available"
	} else {
		Info.Address = taddr.String()
		Info.Address += fmt.Sprintf(":%d", Config.Port)
	}
}

func PrintInfo() {
	fmt.Printf("Name > %s\nVersion > %s\nAddress > %s\nHostname > %s\nContainter > %s\nOS > %s/%s\nDebug > %v\nPath to config > %s\n",
		Config.Name, getVersion(), Info.Address, Info.Hostname,
		Info.Container, Info.OS, Info.Arch, *Debug, *PathToConfig)
}

func GetVersion() string {
	return getVersion()
}

func GetName() string {
	return Config.Name
}

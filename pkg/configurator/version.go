package configurator

var version string

func getVersion() string {
	return version
}

func SetVersion(v string) {
	version = v
}

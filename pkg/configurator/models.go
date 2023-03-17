package configurator

const defaultName = "ErrataID"
const defaultPort = 9111

type ConfigT struct {
	DataBase       string   `yaml:"database"`
	Login          string   `yaml:"login"`
	Password       string   `yaml:"password"`
	AddressToClick string   `yaml:"clickhouse_address"`
	DialTimeout    int      `yaml:"dialTimeout"`
	HTTP           bool     `yaml:"HTTP"`
	Allowed        []string `yaml:"allowed"`
	Name           string   `yaml:"name"`
	Port           uint16   `yaml:"port"`
}

type InfoT struct {
	Name      string
	Hostname  string
	Address   string
	Container string
	OS        string
	Arch      string
}

package configurator

type ConfigT struct {
	Port           string   `yaml:"port"`
	DataBase       string   `yaml:"database"`
	Login          string   `yaml:"login"`
	Password       string   `yaml:"password"`
	AddressToClick string   `yaml:"clickhouse_address"`
	DialTimeout    int      `yaml:"dialTimeout"`
	HTTP           bool     `yaml:"HTTP"`
	Allowed        []string `yaml:"allowed"`
}

type InfoT struct {
	Hostname  string
	Address   string
	Container string
	OS        string
	Arch      string
}

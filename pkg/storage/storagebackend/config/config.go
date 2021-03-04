package config

var DefaultPrefix string = "vehicle-insight"

type TransportConfig struct {
	ServerList    []string
	KeyFile       string
	CertFile      string
	TrustedCAFile string
}

type Config struct {
	Type      string
	Prefix    string
	Transport TransportConfig
}

func NewDefaultConfig() Config {
	return Config{
		Prefix: DefaultPrefix,
	}
}

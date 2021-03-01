package config

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

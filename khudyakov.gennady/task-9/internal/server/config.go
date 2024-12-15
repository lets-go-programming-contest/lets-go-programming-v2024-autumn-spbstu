package server

type ServerConfig struct {
	Host string `yaml:"host" validate:"hostname_port"`
}

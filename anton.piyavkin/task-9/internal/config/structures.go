package config

type Config struct {
	DataBase ConfigDB `yaml:"database"`
	Server   Server   `yaml:"server"`
}

type ConfigDB struct {
	UserDB     string `yaml:"dbUser"`
	PasswordDB string `yaml:"dbPassword"`
	NameDB     string `yaml:"dbName"`
	HostDB     string `yaml:"dbHost"`
	PortDB     int    `yaml:"dbPort"`
}

type Server struct {
	Port string `yaml:"port"`
}

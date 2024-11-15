package flag

import "flag"

var configPath string

func InitFlag() string {
	flag.StringVar(&configPath, "config", "config.yml", "The path to the configuration file")
	flag.Parse()
	return configPath
}

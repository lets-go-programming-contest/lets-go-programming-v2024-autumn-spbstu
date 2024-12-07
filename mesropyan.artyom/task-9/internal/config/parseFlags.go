package config

import "flag"

var fileName string

func init() {
	flag.StringVar(&fileName, "config", "./configs/dbConfig.yaml", "Read file with configuration data")
	flag.StringVar(&fileName, "config", "./configs/serverConfig.yaml", "Read file with configuration data")
	flag.Parse()
}

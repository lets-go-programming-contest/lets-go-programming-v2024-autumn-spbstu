package config

import (
	"flag"
)

func ReadConfigFlag() string {
	cLocation := ""
	flag.StringVar(&cLocation, "config", "", "the config file")
	flag.Parse()
	return string(cLocation)
}

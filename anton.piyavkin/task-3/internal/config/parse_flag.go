package config

import (
	"flag"
)

var NameFile string

func init() {
	flag.StringVar(&NameFile, "config", "default.yaml", "Name file to read from")
	flag.Parse()
}

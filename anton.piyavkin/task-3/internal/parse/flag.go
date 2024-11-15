package parse

import (
	"flag"
)

var NameFile string

func init() {
	flag.StringVar(&NameFile, "config", "", "Name file to read from")
	flag.Parse()
}

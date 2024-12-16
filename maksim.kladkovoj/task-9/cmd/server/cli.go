package main

import "flag"

var ConfigPathFlag = flag.String("config", "config/defauilt.yaml", "Path to YAML config")

func init() {
	flag.Parse()
}

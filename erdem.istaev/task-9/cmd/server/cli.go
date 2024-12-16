package main

import "flag"

var (
	ConfigPathFlag = flag.String("config", "configs/config.yaml", "Path to YAML config")
)

func init() {
	flag.Parse()
}

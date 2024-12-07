package main

import "flag"

var ConfigPathFlag *string

func init() {
	flag.StringVar(ConfigPathFlag, "config", "../../configs/config.yml", "Path to YAML config")
}

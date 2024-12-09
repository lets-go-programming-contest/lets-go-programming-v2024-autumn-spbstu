package config

import "flag"

func GetPath() string {
	var path string
	flag.StringVar(&path, "config", "configs/default.yaml", "config file path")
	flag.Parse()
	return path
}

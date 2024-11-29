package config

import "flag"

func GetPath() string{
	var path string
	flag.StringVar(&path, "config", "config.yml", "config file path")
	flag.Parse()
	return path
}

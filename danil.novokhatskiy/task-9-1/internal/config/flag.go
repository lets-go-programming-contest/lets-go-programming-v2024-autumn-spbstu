package config

import "flag"

func GetPathOfFile() string {
	var path string
	flag.StringVar(&path, "config", "config.yaml", "config file path")
	flag.Parse()
	return path
}

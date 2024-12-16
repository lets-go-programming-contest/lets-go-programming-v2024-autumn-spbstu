package flags

import "flag"

func ConfigParse() string {
	var ConfigPathFlag string
	flag.StringVar(&ConfigPathFlag, "config", "D:\\golang\\lets-go-programming-v2024-autumn-spbstu\\elina.klebleeva\\task-9\\configs\\config.yml", "Path to YAML config")
	flag.Parse()

	return ConfigPathFlag
}

package flags

import "flag"

func ConfigParse() string {
	var ConfigPathFlag string
	flag.StringVar(&ConfigPathFlag, "config", "..\\..\\configs\\config.yml", "Path to YAML config")
	flag.Parse()

	return ConfigPathFlag
}

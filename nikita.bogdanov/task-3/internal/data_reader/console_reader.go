package data_reader

import "flag"

const (
	configFlag        = "config"
	defaultConfigPath = "C:\\Users\\nik_b\\OneDrive\\Desktop\\git_vs\\nikita.bogdanov\\task-3\\config\\config,yaml"
	configDescription = "Configuration path"
)

func ConfigFlag() string {
	var configPath string
	flag.StringVar(&configPath, configFlag, defaultConfigPath, configDescription)
	flag.Parse()
	return configPath
}

package config

import (
	"flag"
	"log"
)

var (
	dbFileName     string
	serverFileName string
	configFilePath string
)

func init() {
	flag.StringVar(&configFilePath, "config", "C:\\Users\\korao\\GolandProjects\\task-9\\internal\\configFiles\\paths.yaml", "Main configuration file with paths")
	flag.Parse()

	configPaths, err := loadConfigPaths(configFilePath)
	if err != nil {
		log.Fatalf("Config file load error: %v", err)
	}

	dbFileName = configPaths.DBConfig
	serverFileName = configPaths.ServerConfig
	log.Printf("Paths from YAML loaded : dbFileName=%s, serverFileName=%s", dbFileName, serverFileName)
}

package main

import (
	"log"

	"github.com/IDevFrye/task-3/internal/config"
)

func main() {
	// Получение пути к файлу конфигурации
	configPath := config.GetConfigPath()

	// Загрузка конфигурации
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
}

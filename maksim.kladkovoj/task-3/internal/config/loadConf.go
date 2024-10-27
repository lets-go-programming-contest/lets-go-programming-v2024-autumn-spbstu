package config

import (
	"fmt"
	"os"

	strct "github.com/Mmmakskl/task-3/internal/structures"
	"gopkg.in/yaml.v3"
)

func LoadConfig(path string, conf *strct.Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Error open files: %w", err)
	}

	if err := yaml.Unmarshal(data, conf); err != nil {
		return fmt.Errorf("Failed decode: %w", err)
	}

	return nil
}

package parseFlag

import (
	"flag"
	"fmt"
	"os"

	"github.com/EmptyInsid/task-3/internal/errorUtils"
)

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return errorUtils.ErrorWithLocation(err)
	}
	if s.IsDir() {
		return errorUtils.ErrorWithLocation(fmt.Errorf("'%s' is a directory", path))
	}
	return nil
}

func ParseFlags() (string, error) {

	var configPath string
	flag.StringVar(&configPath, "config", "./data/config.yml", "path to config file")

	flag.Parse()

	if err := ValidateConfigPath(configPath); err != nil {
		return "", errorUtils.ErrorWithLocation(err)
	}
	return configPath, nil
}

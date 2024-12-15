package app

import (
	"github.com/KRYST4L614/task-9/internal/db"
	"github.com/KRYST4L614/task-9/internal/server"
)

type Config struct {
	ServerConfig server.ServerConfig `yaml:"server"`
	DbConfig     db.DBConfig         `yaml:"db"`
}

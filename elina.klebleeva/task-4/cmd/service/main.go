package main

import (
	"github.com/EmptyInsid/task-4/internal/app"
	"github.com/EmptyInsid/task-4/internal/parseFlag"
)

func main() {
	appVersion := parseFlag.ParseFlags()
	app.Run(appVersion)
}

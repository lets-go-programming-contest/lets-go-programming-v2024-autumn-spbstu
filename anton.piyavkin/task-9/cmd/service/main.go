package main

import (
	"github.com/Piyavva/task-9/internal/config"
	"github.com/Piyavva/task-9/internal/flag"
)

func main() {
	cfg := config.Load(flag.NameFile)

}

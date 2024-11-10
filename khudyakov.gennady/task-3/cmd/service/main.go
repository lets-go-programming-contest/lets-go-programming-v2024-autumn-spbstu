package main

import (
	"bytes"
	"flag"
	"os"
	"sort"

	"github.com/KRYST4L614/task-3/internal/config"
	"github.com/KRYST4L614/task-3/internal/currencies"
	"github.com/KRYST4L614/task-3/internal/io"
)

func main() {
	configFilePath := flag.String("config", "", "config file path")
	flag.Parse()

	configFile, err := os.ReadFile(*configFilePath)
	if err != nil {
		panic(err)
	}

	config, err := (&config.Config{}).GetConfig(configFile)
	if err != nil {
		panic(err)
	}

	source, err := os.ReadFile(config.InputFile)
	if err != nil {
		panic(err)
	}

	source = bytes.ReplaceAll(source, []byte(","), []byte("."))

	currencies, err := (&currencies.Currencies{}).ParseXML(source)
	if err != nil {
		panic(err)
	}

	sort.Sort(&currencies)
	data, err := currencies.ConvertToJSON()
	if err != nil {
		panic(err)
	}
	io.WriteFile(data, config.OutputFile)
}

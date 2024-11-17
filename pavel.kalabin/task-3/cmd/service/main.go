package main

import (
    "flag"

    "github.com/zafod42/task-3/internal/input"
    "github.com/zafod42/task-3/internal/ioconfig"
    "github.com/zafod42/task-3/internal/conversions"

)

func main() {
    var (
        configPath string
        configuration ioconfig.Config
    )

    // Read Flags
    flag.StringVar(&configPath, "config", "", "Path to configuration file")
    flag.Parse()

    err := configuration.ReadFromFile(configPath)
    if err != nil {
        panic(err)
    }

    // Process XML data
    data, err := conversions.ProcessDataFromFile(configuration.InputFile)
    if err != nil {
        panic(err)
    }

    // Write JSON output
    err = input.WriteResult(configuration.OutputFile, data)
    if err != nil {
        panic(err)
    }
}

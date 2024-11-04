package main

import (
    "flag"
    "os"
    "strings"
    "path/filepath"

    "github.com/zafod42/task-3/internal/currenciesJson"
    "github.com/zafod42/task-3/internal/currencies"
    "github.com/zafod42/task-3/internal/ioconfig"

)

func main() {
    var (
        configPath string // what will be default
        configuration ioconfig.Config
        currenciesList currencies.Currencies 
        filteredCurrencies currenciesJson.FilteredCurrencies
    )

    flag.StringVar(&configPath, "config", "", "Path to configuration file")
    flag.Parse()

    contents, err := os.ReadFile(configPath)
    if err != nil {
        panic(err)
    }
    err = configuration.Parse(contents)
    if err != nil {
        panic(err)
    }
    contents, err = os.ReadFile(configuration.InputFile)
    if err != nil {
        panic(err)
    }
    contents = []byte(strings.ReplaceAll(string(contents), ",", "."))
    err = currenciesList.Parse(contents)
    if err != nil {
        panic(err)
    }
    currenciesList.Sort()
    filteredCurrencies.Filter(currenciesList)
    data, err := filteredCurrencies.Dump()
    if err != nil {
        panic(err)
    }
    directory := filepath.Dir(configuration.OutputFile)
    err = os.MkdirAll(directory, os.ModePerm)
    if err != nil {
        panic(err)
    }
    file, err := os.OpenFile(configuration.OutputFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
    defer file.Close()
    if err != nil {
        panic(err)
    }
    os.WriteFile(file.Name(), data, os.ModePerm)
    if err != nil {
        panic(err)
    }
}

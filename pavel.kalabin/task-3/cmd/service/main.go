package main

import (
    "fmt"
    "flag"
    "os"
    "strings"

    "github.com/zafod42/task-3/internal/currencies"
    "github.com/zafod42/task-3/internal/ioconfig"
)

func main() {
    var (
        configPath string // what will be default
        configuration ioconfig.Config
        currenciesList currencies.Currencies 
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
    fmt.Println(currenciesList)
}

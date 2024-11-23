package conversions

import (

    "strings"

    "github.com/zafod42/task-3/internal/input"
    "github.com/zafod42/task-3/internal/currenciesJson"
    "github.com/zafod42/task-3/internal/currencies"
)

func ProcessDataFromFile(FilePath string) (Processed []byte, err error)  {

    var (
        currenciesList currencies.Currencies
        filteredCurrencies currenciesJson.FilteredCurrencies
    )
    contents, err := input.ReadXML(FilePath)
    if err != nil {
        return
    }
    contents = []byte(strings.ReplaceAll(string(contents), ",", "."))
    err = currenciesList.Parse(contents)
    if err != nil {
        return 
    }
    currenciesList.Sort()
    filteredCurrencies.Filter(currenciesList)
    Processed, err = filteredCurrencies.Dump()
    if err != nil {
        return
    }
    return
}


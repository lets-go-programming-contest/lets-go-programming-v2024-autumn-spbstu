package currenciesJson

import (
    "github.com/zafod42/task-3/internal/currencies"
    "encoding/json"
)

type FilteredEntrie struct {
    NumCode  int     `json:"num_code"`
    CharCode string  `json:"char_code"`
    Value    float32 `json:"value"`
}

type FilteredCurrencies struct {
    Entires []FilteredEntrie 
}

func (f *FilteredCurrencies) Filter(currencieList currencies.Currencies) {
    var entrie FilteredEntrie
    for _, currencie := range currencieList.Entries {
        entrie = FilteredEntrie{currencie.NumCode, currencie.CharCode, currencie.Value}
        f.Entires = append(f.Entires, entrie)
    }
}

func (f* FilteredCurrencies) Dump() ([]byte, error) {
    data, err := json.MarshalIndent(f, "", "  ")
    return data, err
}


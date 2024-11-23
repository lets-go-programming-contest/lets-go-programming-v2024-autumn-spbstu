package currency

import "sort"

func ProcessCurrencies(currencies []Currency) []Currency {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
	return currencies
}

package readFile

import (
	"sort"
)

func SortValutes(valutes *ValCurs) {
	sortFunc := func(i, j int) bool {
		return valutes.Valutes[i].Value > valutes.Valutes[j].Value
	}
	sort.Slice(valutes.Valutes, sortFunc)
}

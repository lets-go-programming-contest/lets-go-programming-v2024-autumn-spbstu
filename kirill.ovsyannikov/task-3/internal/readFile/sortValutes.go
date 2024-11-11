package readFile

import (
	"sort"
	structs "task-3/internal/structs"
)

func SortValutes(valutes *structs.ValCurs) {
	sortFunc := func(i, j int) bool {
		return valutes.Valutes[i].Value > valutes.Valutes[j].Value
	}
	sort.Slice(valutes.Valutes, sortFunc)
}

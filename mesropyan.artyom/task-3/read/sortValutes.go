package read

import (
	"sort"

	"github.com/artem6554/task-3/structs"
)

func SortValutes(valutes *structs.ValCurs) {
	sortFunc := func(i, j int) bool {
		return valutes.Valutes[i].Value < valutes.Valutes[j].Value
	}
	sort.Slice(valutes.Valutes, sortFunc)
}

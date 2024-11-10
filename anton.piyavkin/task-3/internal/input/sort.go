package input

import (
	"github.com/Piyavva/task-3/internal/structures"
	"sort"
)

func SortCurrencies(val *structures.Сurrencies) {
	sort.Slice(val.Quotes, func(i, j int) bool {
		return val.Quotes[i].Value < val.Quotes[j].Value
	})
}

package sort_pkg

import (
	"sort"

	"erdem.istaev/task-3/internal/parser"
)

func SortValutes(valutes []parser.Valute) {
	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})
}

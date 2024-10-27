package logic

import (
	"sort"

	strct "github.com/Mmmakskl/task-3/internal/structures"
)

func SortValutes(ValCurs *strct.ValCurs) {
	sort.SliceStable(ValCurs.Valutes, func(i, j int) bool {
		return ValCurs.Valutes[i].Value > ValCurs.Valutes[j].Value
	})
}

package logic

import "sort"

func SortValutes(ValCurs *ValCurs) {
	sort.SliceStable(ValCurs.Valutes, func(i, j int) bool {
		return ValCurs.Valutes[i].Value > ValCurs.Valutes[j].Value
	})
}

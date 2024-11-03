package sort

import (
	"sort"

	"anastasiya.soboleva/task-3/internal/models"
)

func RatesSort(valutes []models.Currency) {
	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})
}

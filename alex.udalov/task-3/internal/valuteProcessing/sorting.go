package valuteProcessing

import (
	"sort"
	"task-3/internal/valuteStrukts"
)

func Sort(data []valuteStrukts.Valute) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Value > data[j].Value
	})
}

package valuteProcessing

import (
	"github.com/hahapathetic/task-3/internal/valuteStructs"
	"sort"
)

func Sort(data []valuteStructs.Valute) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Value > data[j].Value
	})
}

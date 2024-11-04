package sort

import (
	"sort"

	xmltostruct "github.com/Madyarov-Gleb/task-3/internal/XMLtoStruct"
)

type byValue []xmltostruct.Valute

func (a byValue) Len() int {
	return len(a)
}

func (a byValue) Less(i, j int) bool {
	return a[i].Value > a[j].Value
}

func (a byValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Sorting(data []xmltostruct.Valute) {
	sort.Sort(byValue(data))
}

package currencies

import "sort"

type ByValue []Currency

func (a ByValue) Len() int {
	return len(a)
}

func (a ByValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByValue) Less(i, j int) bool {
	return a[i].Value < a[j].Value
}

func SortCurrencies(currencies []Currency) {
	sort.Sort(ByValue(currencies))
}

package internal

import "sort"

type Val []Currency

func (v Val) Len() int {
	return len(v)
}
func (v Val) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v Val) Less(i, j int) bool {
	return v[i].Value < v[j].Value
}

func SortCurrencies(s []Currency) {
	sort.Sort(Val(s))
}

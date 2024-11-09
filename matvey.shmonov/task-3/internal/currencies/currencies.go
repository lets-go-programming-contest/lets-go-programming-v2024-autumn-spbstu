package currencies

import "sort"

type Currency struct {
	NumCode  string  `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value" json:"value"`
}

type Currencies struct {
	List []Currency `xml:"Valute" json:"valute"`
}

type ByValue []Currency

func (a ByValue) Len() int {
	return len(a)
}

func (a ByValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByValue) Less(i, j int) bool { // DESC
	return a[i].Value > a[j].Value
}

func Sort(currencies []Currency) {
	sort.Sort(ByValue(currencies))
}

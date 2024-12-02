package schemas

type ValuteJSONStructure struct {
	NumCode  int     `json:"number-code"`
	CharCode string  `json:"char-code"`
	Value    float64 `json:"value"`
}

package structures

type ValuteInJSON struct {
	NumCode  int     `json:"num-code"`
	CharCode string  `json:"char-code"`
	Value    float64 `json:"value"`
}

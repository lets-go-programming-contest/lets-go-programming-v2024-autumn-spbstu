package feats

type Feat struct {
	Name string
	Exec func(ops ...float64) float64
}

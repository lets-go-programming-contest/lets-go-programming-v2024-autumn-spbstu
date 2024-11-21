package feats

type Feat struct {
	name string
	exec func(ops ...float64) float64
}

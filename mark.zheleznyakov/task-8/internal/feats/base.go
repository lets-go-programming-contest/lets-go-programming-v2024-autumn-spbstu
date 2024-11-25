package feats

var Features = []Feat{
	Feat{
		"add",
		func(ops ...float64) float64 {
			var res float64 = 0
			for _, i := range ops {
				res += i
			}
			return res
		},
	},
}

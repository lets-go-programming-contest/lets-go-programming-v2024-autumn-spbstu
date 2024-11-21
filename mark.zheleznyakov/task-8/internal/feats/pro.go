//go:build pro
// +build pro

package feats

func init() {
	Features = append(Features,
	  Feat{
      "mul",
      func(ops ...float64) float64 {
        var res float64 = 0
        for _, i := range ops {
          res *= i
        }
        return res
      },
    },
	)
}

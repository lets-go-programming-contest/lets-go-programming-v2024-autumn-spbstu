package bounds

import "fmt"

type TempBounds struct {
	LowerBound int
	UpperBound int
}

func (bounds *TempBounds) checkBounds() bool {
	return bounds.LowerBound <= bounds.UpperBound
}

func (bounds *TempBounds) EditBounds(cond string, temp int) {
	if cond == ">=" {
		bounds.LowerBound = max(bounds.LowerBound, temp)
	} else if cond == "<=" {
		bounds.UpperBound = min(bounds.UpperBound, temp)
	}
	if bounds.checkBounds() {
		fmt.Println(bounds.LowerBound)
	} else {
		fmt.Println(-1)
	}
}

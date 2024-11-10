package structure

import (
	"erdem.istaev/task-2-1/internal/comparsion"
)

type TemperatureBounds struct {
	Lower, Upper int
}

func (tb *TemperatureBounds) UpdateBounds(condition string, temp int) {
	if condition == ">=" {
		tb.Lower = comparsion.MaxNum(tb.Lower, temp)
	} else if condition == "<=" {
		tb.Upper = comparsion.MinNum(tb.Upper, temp)
	}
}

func (tb *TemperatureBounds) IsValid() bool {
	return tb.Lower <= tb.Upper
}

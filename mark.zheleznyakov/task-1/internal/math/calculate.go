package math

func Calculate(c *Calculation) float64 {
	switch c.Operator {
	case '+':
		return c.LeftOperand + c.RightOperand
	case '-':
		return c.LeftOperand - c.RightOperand
	case '*':
		return c.LeftOperand * c.RightOperand
	case '/':
		if c.RightOperand == 0 {
			panic("Zero division is illegal")
		}
		return c.LeftOperand / c.RightOperand
	default:
		panic("Idk anything about this operator")
	}
}

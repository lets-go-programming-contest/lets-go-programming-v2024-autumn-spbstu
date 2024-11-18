package unsafe 

type UnsafeCounter struct {
	Val int
}

func (c *UnsafeCounter) Count() {
	c.Val = c.Val + 1
}

func (c *UnsafeCounter) Value() int {
	return c.Val
}


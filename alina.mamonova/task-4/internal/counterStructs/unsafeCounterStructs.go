package counterStructs

type UnsafeCounter struct {
	Count int
}

func (counter *UnsafeCounter) Increment() {
	counter.Count++
}

func (counter *UnsafeCounter) Value() int {
	return counter.Count
}

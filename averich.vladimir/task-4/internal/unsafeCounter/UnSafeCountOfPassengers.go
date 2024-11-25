package unsafeCounter

type TurnstileCounter struct {
	count int
}

func NewTurnstileCounter(count int) *TurnstileCounter {
	return &TurnstileCounter{
		count: count,
	}
}

func (tc *TurnstileCounter) PassThrough() {
	tc.count++
}

func (tc *TurnstileCounter) GetCount() int {
	return tc.count
}
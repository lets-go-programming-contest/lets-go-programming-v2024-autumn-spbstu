package turnstile

type UnsafeTurnstile struct {
	people int
}

func (t *UnsafeTurnstile) Increment() {
	t.people++
}

func (t *UnsafeTurnstile) GetValue() int {
	return t.people
}

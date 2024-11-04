package metro

type UnsafeVisitors struct {
	Quantity map[int]struct{}
}

type UnsafeVisitorRepo interface {
	UnsafeRegister(id int)
	UnsafeGetCount() int
	UnsafeSimulator(cntVisitors int, out chan<- string)
}

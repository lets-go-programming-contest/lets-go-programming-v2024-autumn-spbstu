package mock

//go:generate mockery --all --testonly --quiet --outpkg mock_test --output .

type Doer interface {
	DoSomething(int, string) error
}

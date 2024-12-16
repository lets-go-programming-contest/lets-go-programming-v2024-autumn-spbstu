// go:generate mockery --all --testonly --quiet --outpkg mock_test --output .
package mock

type Doer interface {
	DoSomething(int, string) error
}

// go generate mockgen -destination=mock_foo.go -package=generate . Doer
package mock

type Doer interface {
	DoSomething(int, string) error
}

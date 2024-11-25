package third

//go:generate mockgen -destination=generate_mock.go -package=third . Example
type Example interface {
	Something(a, b int) error
}

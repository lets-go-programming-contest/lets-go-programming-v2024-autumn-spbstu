//go:generate mockgen -destination=mock_service.go -package=generator . Service

package generator

import "fmt"

type Service interface {
	DoSomething() string
}

type RealService struct{}

func (r *RealService) DoSomething() string {
	return "Real Service is working"
}

func main() {
	service := &RealService{}
	fmt.Println(service.DoSomething())
}

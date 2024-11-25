package service

//go:generate mockgen -source=main.go -destination=service.go -package=service

type UserService interface {
	GetUser(id string) (string, error)
}

func main() {
}

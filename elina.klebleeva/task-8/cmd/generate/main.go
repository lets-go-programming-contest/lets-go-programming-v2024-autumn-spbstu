package main

//go:generate mockgen -source=main.go -destination=mocks/mock_database.go -package=mocks

type Database interface {
	GetData(id int) string
	SaveData(id int, data string) error
}

func main() {
	//maybe add
}

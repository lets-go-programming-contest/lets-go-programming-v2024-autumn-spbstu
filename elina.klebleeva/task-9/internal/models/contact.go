package models

type Contact struct {
	Id    int    `json: "id"`
	Name  string `json: "name"`
	Phone string `json: "phone"`
}

package domain

type Contact struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

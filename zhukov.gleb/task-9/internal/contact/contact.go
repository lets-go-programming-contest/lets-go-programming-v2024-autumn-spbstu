package contact

import "errors"

type Contact struct {
	ID    int
	Name  string
	Phone string
}

type ContactInface interface {
	GetAll() ([]Contact, error)
	GetByID(id int) (Contact, error)
	Add(name, phone string) (Contact, error)
	Update(id int, name, phone string) (Contact, error)
	Delete(id int) error
}

var (
	ErrIncorrectPhone = errors.New("error got incorrect phone No")
)

package contact

type Contact struct {
	name  string
	phone string
}

type ContactInface interface {
	GetAll() ([]Contact, error)
	GetByID(id int) (Contact, error)
	Add(contact Contact) error
	Update(id int, contact Contact) error
	Delete(id int) error
}

package contact

type Contact struct {
	name  string
	phone string
}

type ContactInface interface {
	GetAll() ([]Contact, error)
	GetByID(id uint32) (Contact, error)
	Add(contact Contact) error
	Update(id uint32, contact Contact) error
	Delete(id uint32) error
}

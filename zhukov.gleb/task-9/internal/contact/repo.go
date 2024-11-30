package contact

import (
	"fmt"
	"regexp"
)

type ContactDB interface {
	AddContact(name string, phone string) (Contact, error)
	GetContact(id int) (Contact, error)
	DeleteContact(id int) error
	GetAllContacts() ([]Contact, error)
	UpdateContact(id int, name, phone string) (Contact, error)
}

type ContactRepository struct {
	ContactDB
}

func NewContactRepo(repo ContactDB) *ContactRepository {
	return &ContactRepository{repo}
}

func (r *ContactRepository) GetAll() ([]Contact, error) {
	return r.ContactDB.GetAllContacts()
}

func (r *ContactRepository) GetByID(id int) (Contact, error) {
	return r.ContactDB.GetContact(id)
}

var re = regexp.MustCompile("^(\\+7|8)?[\\s\\-]?\\(?[489][0-9]{2}\\)?[\\s\\-]?[0-9]{3}[\\s\\-]?[0-9]{2}[\\s\\-]?[0-9]{2}$")

func validatePhone(phone string) bool {
	return re.MatchString(phone)
}

func (r *ContactRepository) Add(name, phone string) (Contact, error) {
	if !validatePhone(phone) {
		return Contact{}, fmt.Errorf("add: %w", ErrIncorrectPhone)
	}

	return r.ContactDB.AddContact(name, phone)
}

func (r *ContactRepository) Update(id int, name, phone string) (Contact, error) {
	return r.ContactDB.UpdateContact(id, name, phone)
}

func (r *ContactRepository) Delete(id int) error {
	return r.ContactDB.DeleteContact(id)
}

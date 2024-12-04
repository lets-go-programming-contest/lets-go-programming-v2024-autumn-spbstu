package contact

import (
	"fmt"
	"regexp"
	"sort"
	"time"
)

type ContactDBInface interface {
	AddContact(name, phone string, time time.Time) (Contact, error)
	GetContact(id int) (Contact, error)
	DeleteContact(id int) error
	GetAllContacts() ([]Contact, error)
	UpdateContact(id int, name, phone string, time time.Time) (Contact, error)
}

type ContactRepository struct {
	ContactDBInface
}

func NewContactRepo(repo ContactDBInface) *ContactRepository {
	return &ContactRepository{repo}
}

func (r *ContactRepository) GetAll(orderBy string) ([]Contact, error) {
	contacts, err := r.ContactDBInface.GetAllContacts()
	if err != nil {
		return contacts, err
	}

	switch orderBy {
	case "id":
		sort.Slice(contacts, func(i, j int) bool {
			return contacts[i].ID < contacts[j].ID
		})
	case "phone":
		sort.Slice(contacts, func(i, j int) bool {
			return contacts[i].Phone < contacts[j].Phone
		})
	case "created_at":
		sort.Slice(contacts, func(i, j int) bool {
			return contacts[i].CreatedAt.Before(contacts[j].CreatedAt)
		})
	case "updated_at":
		sort.Slice(contacts, func(i, j int) bool {
			return contacts[i].UpdatedAt.Before(contacts[j].UpdatedAt)
		})
	default:
	}

	return contacts, nil
}

func (r *ContactRepository) GetByID(id int) (Contact, error) {
	return r.ContactDBInface.GetContact(id)
}

// TODO rename phoneRegex
var phoneRegex = regexp.MustCompile("^(\\+7|8)?[\\s\\-]?\\(?[489][0-9]{2}\\)?[\\s\\-]?[0-9]{3}[\\s\\-]?[0-9]{2}[\\s\\-]?[0-9]{2}$")

func validatePhone(phone string) bool {
	return phoneRegex.MatchString(phone)
}

func (r *ContactRepository) Add(name, phone string) (Contact, error) {
	currentTime := time.Now()

	if !validatePhone(phone) {
		return Contact{}, fmt.Errorf("add: %w", ErrIncorrectPhone)
	}

	return r.ContactDBInface.AddContact(name, phone, currentTime)
}

func (r *ContactRepository) Update(id int, name, phone string) (Contact, error) {
	currentTime := time.Now()

	if !validatePhone(phone) {
		return Contact{}, fmt.Errorf("add: %w", ErrIncorrectPhone)
	}

	return r.ContactDBInface.UpdateContact(id, name, phone, currentTime)
}

func (r *ContactRepository) Delete(id int) error {
	return r.ContactDBInface.DeleteContact(id)
}

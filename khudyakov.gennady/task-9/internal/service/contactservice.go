package service

import (
	"context"
	"fmt"
	"regexp"

	"github.com/KRYST4L614/task-9/internal/domain"
	"github.com/KRYST4L614/task-9/internal/errlib"
	"github.com/KRYST4L614/task-9/internal/repository"
)

type contactService struct {
	repo repository.ContactRepositoryIface
}

func NewContactService(repo repository.ContactRepositoryIface) *contactService {
	return &contactService{
		repo: repo,
	}
}

func (r *contactService) GetContact(ctx context.Context, id int) (*domain.Contact, error) {
	return r.repo.Get(ctx, id)
}

func (r *contactService) GetAllContacts(ctx context.Context) ([]*domain.Contact, error) {
	return r.repo.GetAll(ctx)
}

func (r *contactService) AddContact(ctx context.Context, contact domain.Contact) (*domain.Contact, error) {
	err := isValidContact(contact)
	if err != nil {
		return nil, err
	}

	return r.repo.Create(ctx, contact)
}

func (r *contactService) UpdateContact(ctx context.Context, contact domain.Contact) (*domain.Contact, error) {
	err := isValidContact(contact)
	if err != nil {
		return nil, err
	}
	return r.repo.Update(ctx, contact)
}

func (r *contactService) DeleteContactById(ctx context.Context, id int) error {
	return r.repo.DeleteById(ctx, id)
}

func isValidContact(contact domain.Contact) error {
	phoneRegex := regexp.MustCompile(`^\+(\d{1,3})\s\(\d{3}\)\s\d{3}-\d{2}-\d{2}$`)
	if !phoneRegex.MatchString(contact.Phone) {
		return fmt.Errorf("%w: invalid phone number", errlib.ErrBadRequest)
	}
	if len(contact.Name) == 0 {
		return fmt.Errorf("%w: invalid name", errlib.ErrBadRequest)
	}
	return nil
}

package service

import (
	"context"

	"github.com/KRYST4L614/task-9/internal/domain"
)

type ContractServiceIface interface {
	GetContact(ctx context.Context, id int) (*domain.Contact, error)
	GetAllContacts(ctx context.Context) ([]*domain.Contact, error)
	AddContact(ctx context.Context, contact domain.Contact) (*domain.Contact, error)
	UpdateContact(ctx context.Context, contact domain.Contact) (*domain.Contact, error)
	DeleteContactById(ctx context.Context, id int) error
}

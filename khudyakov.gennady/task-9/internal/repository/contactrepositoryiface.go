package repository

import (
	"context"

	"github.com/KRYST4L614/task-9/internal/domain"
)

type ContactRepositoryIface interface {
	Create(ctx context.Context, contact domain.Contact) (*domain.Contact, error)
	Get(ctx context.Context, id int) (*domain.Contact, error)
	GetAll(ctx context.Context) ([]*domain.Contact, error)
	Update(ctx context.Context, contact domain.Contact) (*domain.Contact, error)
	DeleteById(ctx context.Context, id int) error
}

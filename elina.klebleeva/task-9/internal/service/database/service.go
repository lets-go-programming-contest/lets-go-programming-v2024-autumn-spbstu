package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	myErr "github.com/EmptyInsid/task-9/internal/errors"
	"github.com/EmptyInsid/task-9/internal/models"
)

type database interface {
	GetContacts(ctx context.Context) ([]models.Contact, error)
	GetContact(ctx context.Context, id int) (*models.Contact, error)
	CreateContact(ctx context.Context, newContact models.Contact) (int, error)
	UpdateContact(ctx context.Context, contact models.Contact) (*models.Contact, error)
	DeleteContact(ctx context.Context, id int) error
}

type DBService struct {
	db     database
	logger *slog.Logger
}

func NewDBService(db database, logger *slog.Logger) DBService {
	return DBService{
		db:     db,
		logger: logger.With(slog.String("component", "service/database")),
	}
}

func (s *DBService) GetContacts() ([]models.Contact, error) {
	contacts, err := s.db.GetContacts(context.Background())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.logger.Warn("get contacts", slog.String("return", "no contacts"))

			return nil, nil
		}

		s.logger.Error("get contacts", slog.String("return", err.Error()))

		return nil, fmt.Errorf("%w", myErr.ErrInternal)
	}

	return contacts, nil
}

func (s *DBService) GetContact(id int) (*models.Contact, error) {
	contact, err := s.db.GetContact(context.Background(), id)
	if err != nil {
		s.logger.Error("get contact", slog.Int("id", id), slog.String("return", err.Error()))

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("[id=%d] %w", id, myErr.ErrNoContact)
		}

		return nil, fmt.Errorf("%w", myErr.ErrInternal)
	}

	return contact, nil
}

func (s *DBService) CreateContact(contact models.Contact) (int, error) {
	id, err := s.db.CreateContact(context.Background(), contact)
	if err != nil {
		s.logger.Error("create contact", slog.Int("id", id), slog.String("return", err.Error()))

		var pqErr *pgconn.PgError
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return 0, fmt.Errorf("[id=%d] %w", contact.ID, myErr.ErrExistContact)
		} else if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("[id=%d] %w", id, myErr.ErrNoContact)
		}

		return 0, fmt.Errorf("[id=%d] %w", id, myErr.ErrInternal)
	}

	return id, nil
}

func (s *DBService) UpdateContact(contact models.Contact) (*models.Contact, error) {
	newContact, err := s.db.UpdateContact(context.Background(), contact)
	if err != nil {
		s.logger.Error("upd contact", slog.Int("id", contact.ID), slog.String("return", err.Error()))

		var pqErr *pgconn.PgError
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return nil, fmt.Errorf("[id=%d] %w", contact.ID, myErr.ErrExistContact)
		} else if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("[id=%d] %w", contact.ID, myErr.ErrNoContact)
		}

		return nil, fmt.Errorf("[id=%d] %w", contact.ID, myErr.ErrInternal)
	}

	return newContact, nil
}

func (s *DBService) DeleteContact(id int) error {
	if err := s.db.DeleteContact(context.Background(), id); err != nil {
		s.logger.Error("delete contact", slog.Int("id", id), slog.String("return", err.Error()))

		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("[id=%d] %w", id, myErr.ErrNoContact)
		}

		return fmt.Errorf("[id=%d] %w", id, myErr.ErrInternal)
	}

	return nil
}

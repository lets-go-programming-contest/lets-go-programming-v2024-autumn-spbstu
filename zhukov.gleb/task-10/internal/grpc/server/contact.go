package server

import (
	"context"
	"errors"

	gen "task-10/gen/proto/contact/v1"
	"task-10/internal/contact"
	"task-10/internal/db"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ContactServer struct {
	gen.UnimplementedContactServiceServer
	Repo      *contact.ContactRepository
	validator *protovalidate.Validator
}

func NewContactServer(repo *contact.ContactRepository, val *protovalidate.Validator) *ContactServer {
	return &ContactServer{
		Repo:      repo,
		validator: val,
	}
}

func (s *ContactServer) GetContactByID(_ context.Context, req *gen.GetContactByIDRequest) (*gen.GetContactByIDResponse, error) {
	var resp gen.GetContactByIDResponse

	data, err := s.Repo.GetByID(int(req.GetId()))
	switch {
	case err == nil:
	case errors.Is(err, db.ErrNoContact):
		return &resp, status.Errorf(codes.NotFound, db.ErrNoContact.Error())
	default:
		return &resp, status.Errorf(codes.Internal, err.Error())
	}

	resp.Contact = &gen.Contact{
		Id:    int32(data.ID),
		Name:  data.Name,
		Phone: data.Phone,
	}

	return &resp, nil
}

func (s *ContactServer) AddContact(_ context.Context, req *gen.AddContactRequest) (*gen.AddContactResponse, error) {
	var resp gen.AddContactResponse

	data, err := s.Repo.Add(req.GetName(), req.GetPhone())
	switch {
	case err == nil:
	case errors.Is(err, contact.ErrIncorrectPhone):
		return &resp, status.Errorf(codes.InvalidArgument, contact.ErrIncorrectPhone.Error())
	default:
		return &resp, status.Errorf(codes.Internal, err.Error())
	}

	resp.Contact = &gen.Contact{
		Id:    int32(data.ID),
		Name:  data.Name,
		Phone: data.Phone,
	}

	return &resp, nil
}

func (s *ContactServer) GetAllContacts(_ context.Context, req *gen.GetAllContactsRequest) (*gen.GetAllContactsResponse, error) {
	var resp gen.GetAllContactsResponse

	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	contacts, err := s.Repo.GetAll(req.GetOrderBy())
	if err != nil {
		return &resp, status.Errorf(codes.Internal, err.Error())
	}

	var genContacts []*gen.Contact
	for _, c := range contacts {
		genContacts = append(genContacts, &gen.Contact{
			Id:    int32(c.ID),
			Name:  c.Name,
			Phone: c.Phone,
		})
	}

	resp.Contacts = genContacts

	return &resp, nil
}

func (s *ContactServer) UpdateContact(_ context.Context, req *gen.UpdateContactRequest) (*gen.UpdateContactResponse, error) {
	var resp gen.UpdateContactResponse

	data, err := s.Repo.Update(int(req.GetId()), req.GetName(), req.GetPhone())
	switch {
	case err == nil:
	case errors.Is(err, contact.ErrIncorrectPhone):
		return &resp, status.Errorf(codes.InvalidArgument, contact.ErrIncorrectPhone.Error())
	case errors.Is(err, db.ErrNoContact):
		return &resp, status.Errorf(codes.NotFound, db.ErrNoContact.Error())
	default:
		return &resp, status.Errorf(codes.Internal, err.Error())
	}

	resp.Contact = &gen.Contact{
		Id:    int32(data.ID),
		Name:  data.Name,
		Phone: data.Phone,
	}

	return &resp, nil
}

func (s *ContactServer) DeleteContact(_ context.Context, req *gen.DeleteContactRequest) (*gen.DeleteContactResponse, error) {
	var resp gen.DeleteContactResponse

	err := s.Repo.Delete(int(req.GetId()))
	switch {
	case err == nil:
	case errors.Is(err, db.ErrNoContact):
		return &resp, status.Errorf(codes.NotFound, db.ErrNoContact.Error())
	default:
		return &resp, status.Errorf(codes.Internal, err.Error())
	}

	resp.Success = true
	return &resp, nil
}

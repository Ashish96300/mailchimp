package user

import (
	"context"
	"errors"
)

type Service interface {
	Register(ctx context.Context, name, email, passwordHash string) (*User, error)
	GetById(ctx context.Context, id int64) (*User, error)
}

type service struct {
	repo repository
}

func NewService(repo repository) Service {
	return &service{repo: repo}
}

func (s *service) Register(
	ctx context.Context,
	name string,
	email string,
	passwordHash string,
) (*User, error) {

	// Rule: email must be unique
	_, err := s.repo.GetByEmail(ctx, email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	user := &User{
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetById(
	ctx context.Context,
	id int64,
) (*User, error) {
	return s.repo.GetById(ctx, id)
}

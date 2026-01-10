package user

import "context"

type Repository interface {
	Create(ctx context.Context, user *User) error

	GetById(ctx context.Context, id int64) (*User, error)

	GetByEmail(ctx context.Context, email string) (*User, error)

	UpdatePassword(
		ctx context.Context,
		id int64,
		passwordHash string,
	) error

	Delete(ctx context.Context, id int64) error
}

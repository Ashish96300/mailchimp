package contact

import "context"

type repository interface {
	Create(ctx context.Context, contact *Contact) error

	GetById(ctx context.Context, id int64) (*Contact, error)

	ListByAudience(
		ctx context.Context,
		audienceID int64,
	) ([]Contact, error)

	GetByEmail(
		ctx context.Context,
		audienceID int64,
		email string,
	) (*Contact, error)

	UpdateStatus(
		ctx context.Context,
		id int64,
		status ContactStatus,
	) error

	Delete(ctx context.Context, id int64) error
}

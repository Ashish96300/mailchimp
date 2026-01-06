package campaigns

import (
	"context"
	"database/sql"
)

type repository interface {
	Create(ctx context.Context, campaign *Campaign) error
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(
	ctx context.Context,
	campaign *Campaign,
) error {

	query := `
		INSERT INTO campaigns (user_id, subject, body, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		campaign.UserId,  
		campaign.Subject,
		campaign.Body,   
		campaign.Status,
	).Scan(
		&campaign.ID,
		&campaign.CreatedAt,
		&campaign.UpdatedAt,
	)
}
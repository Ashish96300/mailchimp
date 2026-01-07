package campaigns

import (
	"context"
	"database/sql"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) repository {
	return &postgresRepository{db: db}
}

// Create inserts a new campaign and populates ID + timestamps
func (r *postgresRepository) Create(
	ctx context.Context,
	campaign *Campaign,
) error {

	query := `
		INSERT INTO campaigns (user_id, subject, body, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at;
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

// GetById fetches a single campaign by ID
func (r *postgresRepository) GetById(
	ctx context.Context,
	id int64,
) (*Campaign, error) {

	query := `
		SELECT id, user_id, subject, body, status, created_at, updated_at
		FROM campaigns
		WHERE id = $1;
	`

	var campaign Campaign

	err := r.db.QueryRowContext(ctx, query, id).
		Scan(
			&campaign.ID,
			&campaign.UserId,
			&campaign.Subject,
			&campaign.Body,
			&campaign.Status,
			&campaign.CreatedAt,
			&campaign.UpdatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &campaign, nil
}

// UpdateStatus updates only the status of a campaign
func (r *postgresRepository) UpdateStatus(
	ctx context.Context,
	id int64,
	status string,
) error {

	query := `
		UPDATE campaigns
		SET status = $2, updated_at = NOW()
		WHERE id = $1;
	`

	_, err := r.db.ExecContext(ctx, query, id, status)
	return err
}

// ListByUser returns all campaigns for a given user
func (r *postgresRepository) ListByUser(
	ctx context.Context,
	userID int64,
) ([]Campaign, error) {

	query := `
		SELECT id, user_id, subject, body, status, created_at, updated_at
		FROM campaigns
		WHERE user_id = $1
		ORDER BY created_at DESC;
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []Campaign

	for rows.Next() {
		var c Campaign

		err := rows.Scan(
			&c.ID,
			&c.UserId,
			&c.Subject,
			&c.Body,
			&c.Status,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		campaigns = append(campaigns, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return campaigns, nil
}

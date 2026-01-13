package contact

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

func (r *postgresRepository) Create(
	ctx context.Context,
	contact *Contact,
) error {

	query := `
		INSERT INTO contacts (audience_id, name, email, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at;
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		contact.AudienceId,
		contact.Name,
		contact.Email,
		contact.Status,
	).Scan(
		&contact.ID,
		&contact.CreatedAt,
		&contact.UpdatedAt,
	)
}

func (r *postgresRepository) GetById(
	ctx context.Context,
	id int64,
) (*Contact, error) {

	query := `
		SELECT id, audience_id, name, email, status, created_at, updated_at
		FROM contacts
		WHERE id = $1;
	`

	var contact Contact

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&contact.ID,
		&contact.AudienceId,
		&contact.Name,
		&contact.Email,
		&contact.Status,
		&contact.CreatedAt,
		&contact.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r *postgresRepository) ListByAudience(
	ctx context.Context,
	audienceID int64,
) ([]Contact, error) {

	query := `
		SELECT id, audience_id, name, email, status, created_at, updated_at
		FROM contacts
		WHERE audience_id = $1
		ORDER BY created_at DESC;
	`

	rows, err := r.db.QueryContext(ctx, query, audienceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []Contact

	for rows.Next() {
		var c Contact

		err := rows.Scan(
			&c.ID,
			&c.AudienceId,
			&c.Name,
			&c.Email,
			&c.Status,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}
func (r *postgresRepository) GetByEmail(
	ctx context.Context,
	audienceID int64,
	email string,
) (*Contact, error) {

	query := `
		SELECT id, audience_id, name, email, status, created_at, updated_at
		FROM contacts
		WHERE audience_id = $1 AND email = $2;
	`

	var contact Contact

	err := r.db.QueryRowContext(ctx, query, audienceID, email).Scan(
		&contact.ID,
		&contact.AudienceId,
		&contact.Name,
		&contact.Email,
		&contact.Status,
		&contact.CreatedAt,
		&contact.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r *postgresRepository) UpdateStatus(
	ctx context.Context,
	id int64,
	status ContactStatus,
) error {

	query := `
		UPDATE contacts
		SET status = $2, updated_at = NOW()
		WHERE id = $1;
	`

	_, err := r.db.ExecContext(ctx, query, id, status)
	return err
}

func (r *postgresRepository) Delete(
	ctx context.Context,
	id int64,
) error {

	query := `
		DELETE FROM contacts
		WHERE id = $1;
	`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}


package user

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
	user *User,
) error {

	query := `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at,
		 updated_at;
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}

func (r *postgresRepository) GetById(
	ctx context.Context,
	id int64,
) (*User, error) {

	query := `
		SELECT id, name, email, password_hash, created_at, updated_at
		FROM users
		WHERE id = $1;
	`

	var user User

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *postgresRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*User, error) {

	query := `
		SELECT id, name, email, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1;
	`

	var user User

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *postgresRepository) UpdatePassword(
	ctx context.Context,
	id int64,
	passwordHash string,
) error {

	query := `
		UPDATE users
		SET password_hash = $2, updated_at = NOW()
		WHERE id = $1;
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		id,
		passwordHash,
	)

	return err
}
func (r *postgresRepository) Delete(
	ctx context.Context,
	id int64,
) error {

	query := `
		DELETE FROM users
		WHERE id = $1;
	`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}


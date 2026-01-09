package audience

import(
	"context"
	"database/sql"
)

type postgresRepository struct{
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) repository{
	return &postgresRepository{db:db}
}

func (r* postgresRepository) Create(
	ctx context.Context,
	audience *Audience,
)error{
		query := `
		INSERT INTO audiences (user_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at;
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		audience.UserId,
		audience.Name,
		audience.Description,
	).Scan(
		&audience.ID,
		&audience.CreatedAt,
		&audience.UpdatedAt,
	)
}

func (r *postgresRepository) GetById(
	ctx context.Context,
	id int64,
) (*Audience, error) {

	query := `
		SELECT id, user_id, name, description, created_at, updated_at
		FROM audiences
		WHERE id = $1;
	`

	var audience Audience

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&audience.ID,
		&audience.UserId,
		&audience.Name,
		&audience.Description,
		&audience.CreatedAt,
		&audience.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &audience, nil
}

func (r *postgresRepository) ListByUser(
	ctx context.Context,
	userID int64,
) ([]Audience, error) {

	query := `
		SELECT id, user_id, name, description, created_at, updated_at
		FROM audiences
		WHERE user_id = $1
		ORDER BY created_at DESC;
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var audiences []Audience

	for rows.Next() {
		var a Audience
		err := rows.Scan(
			&a.ID,
			&a.UserId,
			&a.Name,
			&a.Description,
			&a.CreatedAt,
			&a.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		audiences = append(audiences, a)
	}

	return audiences, nil
}

func (r *postgresRepository) Update(
	ctx context.Context,
	audience *Audience,
) error {

	query := `
		UPDATE audiences
		SET name = $1,
		    description = $2,
		    updated_at = now()
		WHERE id = $3;
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		audience.Name,
		audience.Description,
		audience.ID,
	)

	return err
}

func (r *postgresRepository) Delete(
	ctx context.Context,
	id int64,
) error {

	query := `DELETE FROM audiences WHERE id = $1;`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

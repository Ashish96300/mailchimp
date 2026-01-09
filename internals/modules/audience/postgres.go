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

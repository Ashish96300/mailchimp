package emailjob

import(
	"database/sql"
	"context"
)

type postgresRepository struct{
	db *sql.DB
}

func NewPostgresRepository( db *sql.DB) repository{
	return &postgresRepository{db:db}
}

func (r *postgresRepository) Create(
	ctx context.Context,
	emailjob *EmailJob,
) error {

	query := `
		INSERT INTO email_jobs (
			campaign_id,
			contact_id,
			status,
			retry_count,
			last_error
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at, sent_at;
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		emailjob.CampaignId,
		emailjob.ContactId,
		emailjob.Status,
		emailjob.RetryCount,
		emailjob.LastError,
	).Scan(
		&emailjob.ID,
		&emailjob.CreatedAt,
		&emailjob.UpdatedAt,
		&emailjob.SentAt,
	)
}
func (r *postgresRepository) GetById(
	ctx context.Context,
	id int64,
) (*EmailJob, error) {

	query := `
		SELECT
			id,
			campaign_id,
			contact_id,
			status,
			retry_count,
			last_error,
			created_at,
			updated_at,
			sent_at
		FROM email_jobs
		WHERE id = $1;
	`

	var emailjob EmailJob

	err := r.db.QueryRowContext(ctx, query, id).
		Scan(
			&emailjob.ID,
			&emailjob.CampaignId,
			&emailjob.ContactId,
			&emailjob.Status,
			&emailjob.RetryCount,
			&emailjob.LastError,
			&emailjob.CreatedAt,
			&emailjob.UpdatedAt,
			&emailjob.SentAt,
		)

	if err != nil {
		return nil, err
	}

	return &emailjob, nil
}

func (r *postgresRepository) ListPending(
	ctx context.Context,
	limit int,
) ([]EmailJob, error) {

	query := `
		SELECT
			id,
			campaign_id,
			contact_id,
			status,
			retry_count,
			last_error,
			created_at,
			updated_at,
			sent_at
		FROM email_jobs
		WHERE status = $1
		ORDER BY created_at ASC
		LIMIT $2;
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		EmailPending,
		limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []EmailJob

	for rows.Next() {
		var job EmailJob

		err := rows.Scan(
			&job.ID,
			&job.CampaignId,
			&job.ContactId,
			&job.Status,
			&job.RetryCount,
			&job.LastError,
			&job.CreatedAt,
			&job.UpdatedAt,
			&job.SentAt,
		)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return jobs, nil
}

func (r *postgresRepository) MarkProcessing(
	ctx context.Context,
	id int64,
) error {

	query := `
		UPDATE email_jobs
		SET status = $2, updated_at = NOW()
		WHERE id = $1 AND status = $3;
	`

	res, err := r.db.ExecContext(
		ctx,
		query,
		id,
		EmailProcessing,
		EmailPending,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows // job already taken
	}

	return nil
}

func (r *postgresRepository) MarkSent(
	ctx context.Context,
	id int64,
) error {

	query := `
		UPDATE email_jobs
		SET
			status = $2,
			sent_at = NOW(),
			last_error = NULL,
			updated_at = NOW()
		WHERE id = $1;
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		id,
		EmailSent,
	)

	return err
}

func (r *postgresRepository) MarkFailed(
	ctx context.Context,
	id int64,
	errMsg string,
) error {

	query := `
		UPDATE email_jobs
		SET
			status = $2,
			retry_count = retry_count + 1,
			last_error = $3,
			updated_at = NOW()
		WHERE id = $1;
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		id,
		EmailFailed,
		errMsg,
	)

	return err
}

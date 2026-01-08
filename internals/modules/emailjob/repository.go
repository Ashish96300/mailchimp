package emailjob

import (
	"context"
)

type repository interface {
	Create(ctx context.Context, emailJob *EmailJob) error
	GetById(ctx context.Context, id int64) (*EmailJob, error)

	ListPending(ctx context.Context, limit int) ([]EmailJob, error)

	MarkProcessing(ctx context.Context, id int64) error
	MarkSent(ctx context.Context, id int64) error
	MarkFailed(ctx context.Context, id int64, errMsg string) error
}

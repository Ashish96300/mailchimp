package emailjob

import (
	"context"
	"errors"
)

type Service interface {
	Create(
		ctx context.Context,
		campaignID int64,
		contactID int64,
		toEmail string,
	) (*EmailJob, error)

	GetById(ctx context.Context, id int64) (*EmailJob, error)

	ListPending(ctx context.Context, limit int) ([]EmailJob, error)

	MarkProcessing(ctx context.Context, id int64) error
	MarkSent(ctx context.Context, id int64) error
	MarkFailed(ctx context.Context, id int64, errMsg string) error
}

type service struct {
	repo repository
}

func NewService(repo repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(
	ctx context.Context,
	campaignID int64,
	contactID int64,
	toEmail string,
) (*EmailJob, error) {

	if campaignID <= 0 || contactID <= 0 || toEmail == "" {
		return nil, errors.New("invalid input for creating email job")
	}

	emailJob := &EmailJob{
		CampaignId: campaignID,
		ContactId:  contactID,
		Status:     EmailPending,
		RetryCount: 0,
	}

	if err := s.repo.Create(ctx, emailJob); err != nil {
		return nil, err
	}

	return emailJob, nil
}

func (s *service) GetById(ctx context.Context, id int64) (*EmailJob, error) {
	if id <= 0 {
		return nil, errors.New("invalid email job ID")
	}
	return s.repo.GetById(ctx, id)
}

func (s *service) ListPending(ctx context.Context, limit int) ([]EmailJob, error) {
	if limit <= 0 {
		limit = 10 // default batch size
	}
	return s.repo.ListPending(ctx, limit)
}

func (s *service) MarkProcessing(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid email job ID")
	}
	return s.repo.MarkProcessing(ctx, id)
}

func (s *service) MarkSent(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid email job ID")
	}
	return s.repo.MarkSent(ctx, id)
}

func (s *service) MarkFailed(ctx context.Context, id int64, errMsg string) error {
	if id <= 0 {
		return errors.New("invalid email job ID")
	}
	return s.repo.MarkFailed(ctx, id, errMsg)
}

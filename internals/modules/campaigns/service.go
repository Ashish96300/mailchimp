package campaigns

import (
	"context"
	"errors"
)

type Service interface {
	Register(
		ctx context.Context,
		subject string,
		body string,
		userID int64,
		audienceID int64,
	) (*Campaign, error)

	GetById(ctx context.Context, id int64) (*Campaign, error)
}

type service struct {
	repo repository
}

func NewService(repo repository) Service {
	return &service{repo: repo}
}

func (s *service) Register(
	ctx context.Context,
	subject string,
	body string,
	userID int64,
	audienceID int64,
) (*Campaign, error) {

	// basic validation (same level as user.Register)
	if subject == "" || body == "" {
		return nil, errors.New("subject and body are required")
	}

	if userID <= 0 || audienceID <= 0 {
		return nil, errors.New("invalid user or audience")
	}

	campaign := &Campaign{
		UserId:     userID,
		AudienceId: audienceID,
		Subject:   subject,
		Body:      body,
		Status:    StatusDraft, // default state
	}

	if err := s.repo.Create(ctx, campaign); err != nil {
		return nil, err
	}

	return campaign, nil
}

func (s *service) GetById(
	ctx context.Context,
	id int64,
) (*Campaign, error) {
	return s.repo.GetById(ctx, id)
}

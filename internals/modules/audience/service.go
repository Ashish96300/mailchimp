package audience

import "context"

type Service interface {
	Create(ctx context.Context, name, description string, userID int64) (*Audience, error)
	GetById(ctx context.Context, id int64) (*Audience, error)
}

type service struct {
	repo repository
}

func NewService(repo repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(
	ctx context.Context,
	name string,
	description string,
	userID int64,
) (*Audience, error) {

	audience := &Audience{
		UserId:      userID,
		Name:        name,
		Description: &description,
	}

	if err := s.repo.Create(ctx, audience); err != nil {
		return nil, err
	}

	return audience, nil
}
func (s *service) GetById(
	ctx context.Context,
	id int64,
) (*Audience, error) {
	return s.repo.GetById(ctx, id)
}

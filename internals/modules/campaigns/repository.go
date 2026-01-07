package campaigns

import (
	"context"
)

type repository interface{
	Create(ctx context.Context ,campaign *Campaign) error
	GetById(ctx context.Context ,id int64)(*Campaign ,error)
	UpdateStatus(ctx context.Context ,id int64 ,status string) error
	ListByUser(ctx context.Context ,userID int64) ([]Campaign ,error)
} 
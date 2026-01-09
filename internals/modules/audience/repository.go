package audience

import(
	"context"
)

type repository interface{
	Create(ctx context.Context ,audience *Audience) error
	GetById(ctx context.Context ,id int64)(*Audience ,error)
	ListByUser(ctx context.Context ,userID int64)([]Audience ,error)
	Update(ctx context.Context ,audience *Audience) error
	Delete(ctx context.Context ,id int64)error
}
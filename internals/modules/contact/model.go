package contact

import "time"

type ContactStatus string

const(
	StatusSubscribed       ContactStatus="subscribed"
	StatusUnsubscribed     ContactStatus="unsubscribed"
)

type Contact struct {
	ID           int64     `json:"id" db:"id"`
	AudienceId	 int64      `json:"audience_id" db:"audience_id"` 
	Name         string    `json:"name" db:"name"`         
	Email        string    `json:"email" db:"email"`
	Status		 ContactStatus   `json:"status" db:"status"`


	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

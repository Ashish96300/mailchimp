package audience

import "time"

type Audience struct {
	ID           int64     `json:"id" db:"id"`
	UserId		 int64      `json:"user_id" db:"user_id"` 
	Name         string    `json:"name" db:"name"`          //audience name
	Description  string    `json:"description" db:"description"`

	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

//Campaign is instruction

package campaigns

import "time"

type CampaignStatus string

const (
	StatusDraft     CampaignStatus = "draft"
	StatusScheduled CampaignStatus = "scheduled"
	StatusSent      CampaignStatus = "sent"
	StatusFailed    CampaignStatus = "failed"
)

type Campaign struct {
	ID          int64          `json:"id" db:"id"`
	UserId     int64          `json:"user_id" db:"user_id"`
	AudienceId int64          `json:"audience_id" db:"audience_id"`
	Subject     string         `json:"subject" db:"subject"`
	Body        string         `json:"body" db:"body"`
	Status      CampaignStatus `json:"status" db:"status"`

	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}
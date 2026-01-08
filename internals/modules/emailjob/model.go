//Has ONE email been sent to ONE contact?
package emailjob

import "time"

type EmailStatus string

const (
	EmailPending    EmailStatus = "pending"
	EmailProcessing EmailStatus = "processing"
	EmailSent       EmailStatus = "sent"
	EmailFailed     EmailStatus = "failed"
)

type EmailJob struct {
	ID         int64       `json:"id" db:"id"`
	CampaignId int64       `json:"campaign_id" db:"campaign_id"`
	ContactId  int64       `json:"contact_id" db:"contact_id"`

	Status     EmailStatus `json:"status" db:"status"`
	RetryCount int         `json:"retry_count" db:"retry_count"`
	LastError  *string     `json:"last_error" db:"last_error"`

	CreatedAt  time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at" db:"updated_at"`
	SentAt     *time.Time  `json:"sent_at" db:"sent_at"`
}

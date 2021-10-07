package core_domain_model

import (
	"time"
)

type Audit struct {
	Deleted       bool			`bson:"deleted" pg:",soft_delete"`
	CreatedAt     interface{}   `bson:"created_at"`
	UserCreated   interface{}	`bson:"user_created"`
	UpdatedAt     interface{}   `bson:"updated_at"`
	UserUpdated   interface{}	`bson:"user_updated"`
}

func NewAudit(userId interface{}) Audit  {
	audit := Audit{}
	audit.Deleted = false
	audit.UserCreated = userId
	audit.CreatedAt = time.Time{}
	audit.UserUpdated = nil
	audit.UpdatedAt = nil
	return audit
}

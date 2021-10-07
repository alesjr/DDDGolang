package users_domain_model

import "time"

type Users struct {
	Id 		      int 			`bson:"_id"   form:"id"`
	Name 		  string        `bson:"name"  form:"name"  binding:"required"`
	Deleted       bool			`pg:",soft_delete"`
	CreatedAt     time.Time
	UserCreated   int			`pg:",use_zero"`
	UpdatedAt     time.Time
	UserUpdated   int			`pg:",use_zero"`
}
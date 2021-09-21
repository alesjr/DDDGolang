package users_domain_model

import (
	coreDomainModel "ddd_golang/src/Core/Domain/Model"
)

type Users struct {
	Id 		    int 					`bson:"_id"   form:"id"`
	Name 		string             		`bson:"name"  form:"name"  binding:"required"`
	Email 		string             		`bson:"email" form:"email" binding:"required"`
	Audit       coreDomainModel.Audit
}
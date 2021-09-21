package config_users

import (
	"github.com/gin-gonic/gin"
	_userController "ddd_golang/src/Users/Infrastructure/Controller"
)

func NewRoutesUsers(r *gin.Engine, db interface{}) {
	uc := _userController.NewControllerUser(r, db)
	uc.ControllerHttpRoutes()
}
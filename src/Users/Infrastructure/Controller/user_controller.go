package user_controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_coreDomainService "ddd_golang/src/Core/Domain/Service"
	_coreInfrastructureHelper "ddd_golang/src/Core/Infrastructure/Helper"
	_usersUseCaseCreateUser "ddd_golang/src/Users/Application/UseCase/CreateUser"
	_usersDomainModel "ddd_golang/src/Users/Domain/Model"
	_usersInfrastructureRepository "ddd_golang/src/Users/Infrastructure/Repository"
)

type controllerUser struct {
	r *gin.Engine
	db *sql.DB
}

func NewControllerUser(r *gin.Engine, db interface{}) _coreDomainService.ControllerInterface {
	return &controllerUser{
		r,
		db.(*sql.DB),
	}
}

func (c controllerUser) ControllerHttpRoutes() {
	users := *c.r.Group("/users")
	{
		users.GET("", c.getUser)
		users.POST("", c.addUser)
	}
}

func (c controllerUser) getUser(gc *gin.Context) {
	gc.JSONP(200, gin.H{
		"message":"test",
	})
}

func (c controllerUser) addUser(gc *gin.Context) {
	ur := _usersInfrastructureRepository.NewUsersRepository(c.db)

	u := new(_usersDomainModel.Users)
	err := gc.Bind(u)
	if err != nil {
		_coreInfrastructureHelper.GinFormatterErrorRequest(err.(validator.ValidationErrors), gc)
		return
	}

	response, err := _usersUseCaseCreateUser.NewCreateUser(ur).Handle(*u)
	if err != nil {
		_coreInfrastructureHelper.GinFormatterErrorResponse(
			response,
			err,
			gc,
		)
		return
	}

	gc.JSONP(response.GetStatus(), gin.H{
		"data": response.GetData(),
		"message": response.GetMessage(),
	})
}
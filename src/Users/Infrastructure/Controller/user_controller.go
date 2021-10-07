package user_controller

import (
	_coreDomainService "ddd_golang/src/Core/Domain/Service"
	_configCore "ddd_golang/src/Core/Infrastructure/Config"
	_coreInfrastructureHelper "ddd_golang/src/Core/Infrastructure/Helper"
	_usersUseCaseCreateUser "ddd_golang/src/Users/Application/UseCase/CreateUser"
	_usersDomainModel "ddd_golang/src/Users/Domain/Model"
	_usersInfrastructureRepository "ddd_golang/src/Users/Infrastructure/Repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userController struct {
	config *_configCore.Config
}

func NewUserController(config *_configCore.Config) _coreDomainService.ControllerInterface {
	uc := &userController{
		config: config,
	}
	users := config.GetEngine().Group("/users")
	{
		users.GET("", uc.getUser)
		users.POST("", uc.addUser)
	}
	return uc
}

func (c userController) getUser(gc *gin.Context) {
	gc.JSONP(200, gin.H{
		"message":"test",
	})
}

func (c userController) addUser(gc *gin.Context) {
	ur := _usersInfrastructureRepository.NewUsersRepository(c.config.GetDB())

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
package core_controller

import (
	_coreDomainService "ddd_golang/src/Core/Domain/Service"
	_configCore "ddd_golang/src/Core/Infrastructure/Config"
	"github.com/gin-gonic/gin"
)

type defaultController struct {
	config *_configCore.Config
}

func NewDefaultController(config *_configCore.Config) _coreDomainService.ControllerInterface {
	dc := &defaultController{
		config: config,
	}
	users := config.GetEngine().Group("/")
	{
		users.GET("", dc.getStatus)
	}
	return dc
}

func (c defaultController) getStatus(gc *gin.Context) {
	gc.JSONP(200, gin.H{
		"message":"Application Work and is running",
	})
}
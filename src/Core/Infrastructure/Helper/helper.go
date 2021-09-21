package core_infrastructure_helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_coreInfrastructureUsecase "ddd_golang/src/Core/Infrastructure/UseCase"
)

func GinFormatterErrorRequest(ve validator.ValidationErrors, gc *gin.Context) {
	errs := make(map[string]interface{})
	for i := 0; i < len(ve); i++ {
		fe := ve[i]
		errs[fe.Field()] = fmt.Sprintf("field is %s, consult doc for fields params", fe.Tag())
	}
	gc.JSONP(400, gin.H{
		"error": errs,
	})
}

func GinFormatterErrorResponse(response _coreInfrastructureUsecase.UseCaseResponseInterface, err error, gc *gin.Context) {
	gc.JSONP(response.GetStatus(), gin.H{
		"message": response.GetMessage(),
		"error": err.Error(),
	})
}

func GinFormatterSuccessResponse(response _coreInfrastructureUsecase.UseCaseResponseInterface, gc *gin.Context) {
	gc.JSONP(response.GetStatus(), gin.H{
		"data": response.GetData(),
		"message": response.GetMessage(),
	})
}


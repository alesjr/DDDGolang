package config_core

import (
	_coreDomainModel "ddd_golang/src/Core/Domain/Model"
	_coreInfrastructureService "ddd_golang/src/Core/Infrastructure/Service"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func Bootstrap() *Config{
	df := _coreInfrastructureService.NewDatabase(_coreDomainModel.Postgres)
	db := df.Create()
	return newConfig(db.(*pg.DB), gin.Default())
}

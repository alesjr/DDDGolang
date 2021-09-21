package core_config

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_coreDomainModel "ddd_golang/src/Core/Domain/Model"
	_coreInfrastructureService "ddd_golang/src/Core/Infrastructure/Service"
	_configUsers "ddd_golang/src/Users/Config"
)

var db *sql.DB

func Bootstrap() {
	df := _coreInfrastructureService.NewDatabase(_coreDomainModel.Postgres)
	db := df.Create()

	r := gin.Default()
	_configUsers.NewRoutesUsers(r, db)

	r.Run()
}


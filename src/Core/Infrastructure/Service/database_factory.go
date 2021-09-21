package core_infrastructure_service

import (
	"context"
	_ "github.com/jackc/pgx/v4"
	"log"
	_coreDomainModel "ddd_golang/src/Core/Domain/Model"
	_coreDomainService "ddd_golang/src/Core/Domain/Service"
	"time"
)

type databaseFactory struct {
	con interface{}
	ad  _coreDomainService.DatabaseAdapterInterface
}

func NewDatabase(drive string) _coreDomainService.DatabaseFactoryInterface {
	df := &databaseFactory{}

	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Minute)
	if drive == _coreDomainModel.Mongo {
		df.ad = NewMongoDatabaseAdapter(
			"root",
			"root",
			"mongo.local",
			"27017",
			"survey_audit",
			ctx,
		)
	} else if drive == _coreDomainModel.Postgres {
		df.ad = NewPostgresDatabaseAdapter(
			"postgres",
			"1234",
			"postgresql.local",
			"5432",
			"survey",
			ctx,
		)
	} else {
		log.Fatal()
	}
	return df
}

func (d databaseFactory) Create() interface{} {
	if d.con != nil {
		return d.con
	}
	con := d.ad.GetDatabase()
	d.con = con
	return con
}
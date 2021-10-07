package core_infrastructure_service

import (
	"context"
	_coreDomainService "ddd_golang/src/Core/Domain/Service"
	"github.com/go-pg/pg/v10"
	"log"
)

type postgresDatabaseAdapter struct {
	username string
	password string
	host     string
	port     string
	dbName   string
	context  context.Context
	db   *pg.DB
}

func NewPostgresDatabaseAdapter(u string, p string, h string, po string, dn string, ctx context.Context) _coreDomainService.DatabaseAdapterInterface {
	return &postgresDatabaseAdapter{
		username: u,
		password: p,
		host:     h,
		port:     po,
		dbName:   dn,
		context:  ctx,
		db:   nil,
	}
}

func (p *postgresDatabaseAdapter) GetDatabase() interface{} {
	p.conn()
	return p.db
}

func (p *postgresDatabaseAdapter) GetContext() context.Context {
	return p.context
}

func (p *postgresDatabaseAdapter) conn() {
	if p.db == nil {
		cli := pg.Connect(&pg.Options{
			Addr: p.host+":"+p.port,
			User: p.username,
			Password: p.password,
			Database: p.dbName,
		})
		p.db = cli
	}
	err := p.db.Ping(p.GetContext())
	if err != nil {
		log.Fatalln(_coreDomainService.ErrConnectionDB, "pgx", err)
	}
}
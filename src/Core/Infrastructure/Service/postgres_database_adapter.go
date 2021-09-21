package core_infrastructure_service

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	_coreDomainService "ddd_golang/src/Core/Domain/Service"
)

type postgresDatabaseAdapter struct {
	username string
	password string
	host     string
	port     string
	dbName   string
	context  context.Context
	db   *sql.DB
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

func (p *postgresDatabaseAdapter) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", p.username, p.password, p.host, p.port, p.dbName)
}

func (p *postgresDatabaseAdapter) GetDriveName() string {
	return "pgx"
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
		cli, err := sql.Open(p.GetDriveName(), p.GetDSN())
		if err != nil {
			log.Fatalln(_coreDomainService.ErrConnectionDB, p.GetDriveName(), err)
		}
		p.db = cli
	}
	err := p.db.Ping()
	if err != nil {
		log.Fatalln(_coreDomainService.ErrConnectionDB, p.GetDriveName(), err)
	}
}
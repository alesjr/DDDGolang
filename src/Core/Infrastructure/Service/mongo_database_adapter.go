package core_infrastructure_service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	_coreDomainService "ddd_golang/src/Core/Domain/Service"
)

type mongoDatabaseAdapter struct {
	username string
	password string
	host     string
	port     string
	dbName   string
	context  context.Context
	db		 *mongo.Client
}

func NewMongoDatabaseAdapter(u string, p string, h string, po string, dn string, ctx context.Context) _coreDomainService.DatabaseAdapterInterface {
	return &mongoDatabaseAdapter{
		username: u,
		password: p,
		host:     h,
		port:     po,
		dbName:   dn,
		context:  ctx,
		db:   	  nil,
	}
}

func (m *mongoDatabaseAdapter) GetDatabase() interface{} {
	m.conn()
	return m.db.Database(m.dbName)
}

func (m *mongoDatabaseAdapter) GetContext() context.Context {
	return m.context
}

func (m *mongoDatabaseAdapter) conn() {
	if m.db == nil {
		auth := options.Credential{
			Username: m.username,
			Password: m.password,
		}
		oc := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", m.host, m.port)).SetAuth(auth)
		cli, err := mongo.Connect(m.context, oc)
		if err != nil {
			log.Fatalln(_coreDomainService.ErrConnectionDB, "mongo", err)
		}
		m.db = cli
	}
	err := m.db.Ping(m.context, nil)
	if err != nil {
		m.db.Connect(m.context)
	}
}
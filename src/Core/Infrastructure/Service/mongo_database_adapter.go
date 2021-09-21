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
	client   *mongo.Client
}

func NewMongoDatabaseAdapter(u string, p string, h string, po string, dn string, ctx context.Context) _coreDomainService.DatabaseAdapterInterface {
	return &mongoDatabaseAdapter{
		username: u,
		password: p,
		host:     h,
		port:     po,
		dbName:   dn,
		context:  ctx,
		client:   nil,
	}
}

func (m *mongoDatabaseAdapter) GetDSN() string {
	return fmt.Sprintf("mongodb://%s:%s", m.host, m.port)
}

func (m *mongoDatabaseAdapter) GetDriveName() string {
	return "mongo"
}

func (m *mongoDatabaseAdapter) GetDatabase() interface{} {
	m.conn()
	return m.client.Database(m.dbName)
}

func (m *mongoDatabaseAdapter) GetContext() context.Context {
	return m.context
}

func (m *mongoDatabaseAdapter) conn() {
	if m.client == nil {
		auth := options.Credential{
			Username: m.username,
			Password: m.password,
		}
		oc := options.Client().ApplyURI(m.GetDSN()).SetAuth(auth)
		cli, err := mongo.Connect(m.context, oc)
		if err != nil {
			log.Fatalln(_coreDomainService.ErrConnectionDB, m.GetDriveName(), err)
		}
		m.client = cli
	}
	err := m.client.Ping(m.context, nil)
	if err != nil {
		m.client.Connect(m.context)
	}
}
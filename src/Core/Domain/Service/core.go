package core_domain_service

import (
	"context"
	"fmt"
)

var (
	ErrNilConfig      = fmt.Errorf("no config")
	ErrNoDatabaseName = fmt.Errorf("no database name")
	ErrConnectionDB = fmt.Errorf("error connecting to database")
)

type DatabaseFactoryInterface interface {
	Create() interface{}
}

type DatabaseAdapterInterface interface {
	GetDatabase() interface{}
	GetContext() context.Context
}

type ControllerInterface interface {

}
package core_infrastructure_usecase

type UseCaseInterface interface {
	Handle(structRequest interface{}) (UseCaseResponseInterface, error)
}

type UseCaseResponseInterface interface {
	GetData() interface{}
	GetStatus() int
	GetMessage() string
}
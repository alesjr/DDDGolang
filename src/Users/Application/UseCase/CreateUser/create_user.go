package users_usecase_create_user

import (
	_coreInfrastructureUseCase "ddd_golang/src/Core/Infrastructure/UseCase"
	_usersDomainModel "ddd_golang/src/Users/Domain/Model"
	_usersDomainRepository "ddd_golang/src/Users/Domain/Repository"
)

type createUser struct {
	ur _usersDomainRepository.UsersRepositoryInterface
}

type userResponse struct {
	data _usersDomainModel.Users
	status int
	message string
}

func NewCreateUser(ur _usersDomainRepository.UsersRepositoryInterface) _coreInfrastructureUseCase.UseCaseInterface {
	return &createUser{
		ur,
	}
}

func (c createUser) Handle(structRequest interface{}) (_coreInfrastructureUseCase.UseCaseResponseInterface, error){
	u := structRequest.(_usersDomainModel.Users)
	return &userResponse{
		u,
		201,
		"user_create_with_success",
	}, nil
}

func (u userResponse) GetData() interface{} {
	return u.data
}

func (u userResponse) GetStatus() int {
	return u.status
}

func (u userResponse) GetMessage() string {
	return u.message
}
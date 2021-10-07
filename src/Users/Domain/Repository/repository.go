package users_domain_repository

import _usersDomainModel "ddd_golang/src/Users/Domain/Model"

type UsersRepositoryInterface interface {
	Load(id int) (_usersDomainModel.Users, error)
	AddUser(u _usersDomainModel.Users) (_usersDomainModel.Users, error)
	EditUser(u _usersDomainModel.Users) error
	RemoveUser(id int) error
}
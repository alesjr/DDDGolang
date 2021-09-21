package users_domain_repository

import _usersDomainModel "ddd_golang/src/Users/Domain/Model"

type UsersRepositoryInterface interface {
	Load(id int)
	AddUser(u _usersDomainModel.Users)
	EditUser(u _usersDomainModel.Users)
	RemoveUser(id int)
}
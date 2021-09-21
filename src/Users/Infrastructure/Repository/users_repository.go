package users_infrastructure_repository

import (
	"database/sql"
	_usersDomainModel "ddd_golang/src/Users/Domain/Model"
	_usersDomainRepository "ddd_golang/src/Users/Domain/Repository"
)

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) _usersDomainRepository.UsersRepositoryInterface {
	return &usersRepository{
		db: db,
	}
}

func (ur usersRepository) Load(id int) {
	panic("implement me")
}

func (ur usersRepository) AddUser(u _usersDomainModel.Users) {
	panic("implement me")
}

func (ur usersRepository) EditUser(u _usersDomainModel.Users) {
	panic("implement me")
}

func (ur usersRepository) RemoveUser(id int) {
	panic("implement me")
}
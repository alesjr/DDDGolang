package users_infrastructure_repository

import (
	_usersDomainModel "ddd_golang/src/Users/Domain/Model"
	_usersDomainRepository "ddd_golang/src/Users/Domain/Repository"
	"github.com/go-pg/pg/v10"
	_ "github.com/go-pg/pg/v10/orm"
)

type usersRepository struct {
	db *pg.DB
}

func NewUsersRepository(db *pg.DB) _usersDomainRepository.UsersRepositoryInterface {
	return &usersRepository{
		db: db,
	}
}

func (ur usersRepository) Load(id int) (_usersDomainModel.Users, error) {
	panic("implement me")
}

func (ur usersRepository) AddUser(u _usersDomainModel.Users) (_usersDomainModel.Users, error){
	ur.db.Exec("SET search_path TO svy")
	_, err := ur.db.Model(&u).Insert()
	if err != nil {
		return _usersDomainModel.Users{}, err
	}
	return u, nil
}

func (ur usersRepository) EditUser(u _usersDomainModel.Users) error {
	panic("implement me")
}

func (ur usersRepository) RemoveUser(id int) error {
	panic("implement me")
}
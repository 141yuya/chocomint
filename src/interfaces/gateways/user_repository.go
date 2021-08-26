package gateways

import (
	"github.com/141yuya/chocomint/src/domain/entities"
	"github.com/141yuya/chocomint/src/infrastructure"
)

type UserRepository struct {
	infrastructure.SqlHandler
}

func (repo *UserRepository) Persist(u entities.User) (*entities.User, error) {
	user := entities.User{}
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindById(id int) (*entities.User, error) {
	user := entities.User{}
	err := repo.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(handler *infrastructure.SqlHandler) *UserRepository {
	return &UserRepository{SqlHandler: *handler}
}

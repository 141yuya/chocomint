package gateways

import (
	"github.com/141yuya/chocomint/src/domain/entities"
	"github.com/141yuya/chocomint/src/infrastructure"
)

type UserRepository struct {
	infrastructure.SqlHandler
}

func NewUserRepository(handler *infrastructure.SqlHandler) *UserRepository {
	return &UserRepository{SqlHandler: *handler}
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

func (repo *UserRepository) FindAll() (*entities.Users, error) {
	users := entities.Users{}
	err := repo.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (repo *UserRepository) Update(id int, u entities.User) (*entities.User, error) {
	user := entities.User{}
	err := repo.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	saveErr := repo.DB.Save(&user).Error
	if saveErr != nil {
		return nil, saveErr
	}
	return &user, nil
}

func (repo *UserRepository) Delete(id int) error {
	err := repo.DB.Delete(&entities.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

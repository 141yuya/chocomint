package usecases

import (
	"github.com/141yuya/chocomint/src/domain/entities"
	"github.com/141yuya/chocomint/src/domain/repositories"
)

type UserUsecase interface {
	Add(user entities.User) error
}

type userUsecase struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (userUsecase *userUsecase) Add(user entities.User) (err error) {
	_, err = userUsecase.userRepository.Persist(user)
	return err
}

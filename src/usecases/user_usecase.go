package usecases

import (
	"github.com/141yuya/chocomint/src/domain/entities"
	"github.com/141yuya/chocomint/src/domain/repositories"
)

type UserUsecase struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return UserUsecase{userRepository: userRepository}
}

func (userUsecase *UserUsecase) Add(user entities.User) (err error) {
	_, err = userUsecase.userRepository.Persist(user)
	return err
}

func (userUsecase *UserUsecase) GetUser(id int) (*entities.User, error) {
	user, err := userUsecase.userRepository.FindById(id)
	return user, err
}

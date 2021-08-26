package repositories

import "github.com/141yuya/chocomint/src/domain/entities"

type UserRepository interface {
	Persist(entities.User) (*entities.User, error)
	FindById(id int) (*entities.User, error)
}

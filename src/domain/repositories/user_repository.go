package repositories

import "github.com/141yuya/go-clean-architecture/src/domain/entities"

type UserRepository interface {
	Persist(entities.User) (*entities.User, error)
	FindById(id int) (*entities.User, error)
	FindAll() (*entities.Users, error)
	Update(int, entities.User) (*entities.User, error)
	Delete(id int) error
}

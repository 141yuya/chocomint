package usecases_test

import (
	"testing"

	"github.com/141yuya/go-clean-architecture/domain/entities"
	"github.com/141yuya/go-clean-architecture/interfaces/gateways/mock_repositories"
	"github.com/141yuya/go-clean-architecture/usecases"
	"github.com/golang/mock/gomock"
)

func TestAdd(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	var expected *entities.User = &entities.User{ID: 1, FirstName: "First", LastName: "Last"}
	mockUserRepository := mock_repositories.NewMockUserRepository(controller)
	mockUserRepository.EXPECT().Persist(expected).Return(expected, nil)

	userUsecase := usecases.NewUserUsecase(mockUserRepository)
	_, err := userUsecase.Add(expected)

	if err != nil {
		t.Error(err)
	}
}

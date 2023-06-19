package internal

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/regimentor/currency-calc/internal/models"
	"github.com/regimentor/currency-calc/mocks"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStorage := mock_internal.NewMockUserStorage(ctrl)
	userRepository := NewUserRepository(userStorage)

	t.Run("Create user with non empty api key", func(t *testing.T) {
		apiKey := GenerateApiKey()
		dto := models.CreateUserDto{ApiKey: apiKey}
		exceptedUser := &models.User{ApiKey: apiKey, ID: 1}

		userStorage.
			EXPECT().
			Create(context.Background(), dto).
			Return(exceptedUser, nil).
			Times(1)

		_, err := userRepository.Create(
			context.Background(), dto)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("Create user with empty api key", func(t *testing.T) {
		dto := models.CreateUserDto{}

		userStorage.
			EXPECT().
			Create(context.Background(), dto).
			Return(nil, errors.New("create user due err")).
			Times(1)

		_, err := userRepository.Create(
			context.Background(), dto)

		if err == nil {
			t.Errorf("expected error: %v", err)
		}
	})
}

func TestUserRepository_GetByApiKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStorage := mock_internal.NewMockUserStorage(ctrl)
	userRepository := NewUserRepository(userStorage)

	t.Run("Get user by api key", func(t *testing.T) {
		apiKey := GenerateApiKey()
		exceptedUser := &models.User{ApiKey: apiKey, ID: 1}

		userStorage.
			EXPECT().
			GetByApiKey(context.Background(), apiKey).
			Return(exceptedUser, nil).
			Times(1)

		_, err := userRepository.GetByApiKey(
			context.Background(), apiKey)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("Get user by empty api key", func(t *testing.T) {
		apiKey := models.ApiKey("")

		userStorage.
			EXPECT().
			GetByApiKey(context.Background(), apiKey).
			Return(nil, errors.New("get user due err")).
			Times(1)

		_, err := userRepository.GetByApiKey(
			context.Background(), apiKey)

		if err == nil {
			t.Errorf("expected error: %v", err)
		}
	})

}

func TestUserRepository_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStorage := mock_internal.NewMockUserStorage(ctrl)
	userRepository := NewUserRepository(userStorage)

	t.Run("Get user by id", func(t *testing.T) {
		id := models.UserId(1)
		exceptedUser := &models.User{ApiKey: GenerateApiKey(), ID: id}

		userStorage.
			EXPECT().
			GetById(context.Background(), id).
			Return(exceptedUser, nil).
			Times(1)

		_, err := userRepository.GetById(
			context.Background(), id)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("Get user by wrong id", func(t *testing.T) {
		id := models.UserId(0)

		userStorage.
			EXPECT().
			GetById(context.Background(), id).
			Return(nil, errors.New("get user due err")).
			Times(1)

		_, err := userRepository.GetById(
			context.Background(), id)

		if err == nil {
			t.Errorf("expected error: %v", err)
		}

	})
}

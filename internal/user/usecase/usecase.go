package usecase

import (
	"context"

	"go_project_template/internal/user/model"
	"go_project_template/internal/user/repository"
)

type IUserUseCase interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserById(ctx context.Context, id int64) (model.User, error)
	AddNewUser(ctx context.Context, newUser model.User) (int64, error)
	UpdateUser(ctx context.Context, user model.User) (int64, error)
	DeleteUser(ctx context.Context, id int64) error
}

type UserUseCase struct {
	userRepo repository.IUserRepository
}

func NewUserUseCae(userRepo repository.IUserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) GetUsers(ctx context.Context) ([]model.User, error) {
	return nil, nil
}

func (uc *UserUseCase) GetUserById(ctx context.Context, id int64) (model.User, error) {

	userDetails, err := uc.userRepo.GetUserById(ctx, id)

	if err != nil {
		return userDetails, err
	}

	return userDetails, nil
}

func (uc *UserUseCase) AddNewUser(ctx context.Context, newUser model.User) (int64, error) {
	return 0, nil
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, user model.User) (int64, error) {
	return 0, nil
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id int64) error {
	return nil
}

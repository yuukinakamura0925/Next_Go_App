package usecases

import (
	"Next_Go_App/internal/repositories"
	"context"

	"Next_Go_App/ent"
)

type UserUsecase interface {
	SignUp(ctx context.Context, email, name, password string) (*ent.User, error)
	SignIn(ctx context.Context, email, password string) (*ent.User, error)
}

type userUsecase struct {
	userRepo repositories.IUserRepository
}

func NewUserUsecase(userRepo repositories.IUserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) SignUp(ctx context.Context, email, name, password string) (*ent.User, error) {
	user := &ent.User{
		Email:    email,
		Name:     name,
		Password: password,
	}
	return u.userRepo.CreateUser(ctx, user.Email, user.Name, user.Password)
}

func (u *userUsecase) SignIn(ctx context.Context, email, password string) (*ent.User, error) {
	return u.userRepo.GetUserByEmailAndPassword(ctx, email, password)
}

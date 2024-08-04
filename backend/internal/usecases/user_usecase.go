package usecases

import (
	"context"
	"errors"

	"Next_Go_App/ent"
	"Next_Go_App/internal/repositories"
	"Next_Go_App/internal/services"
	"Next_Go_App/internal/utils"
)

type UserUsecase interface {
	SignUp(ctx context.Context, email, name, password string) (*ent.User, error)
	SignIn(ctx context.Context, email, password string) (string, error)
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
}

type userUsecase struct {
	userRepo    repositories.IUserRepository
	passwordSvc services.PasswordService
}

func NewUserUsecase(userRepo repositories.IUserRepository, passwordSvc services.PasswordService) UserUsecase {
	return &userUsecase{
		userRepo:    userRepo,
		passwordSvc: passwordSvc,
	}
}

func (u *userUsecase) SignUp(ctx context.Context, email, name, password string) (*ent.User, error) {
	hashedPassword, err := u.passwordSvc.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return u.userRepo.CreateUser(ctx, email, name, hashedPassword)
}

func (u *userUsecase) SignIn(ctx context.Context, email, password string) (string, error) {
	// ユーザーをメールアドレスで取得
	user, err := u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}
	// パスワードサービスのチェック
	if u.passwordSvc == nil {
		return "", errors.New("password service is not initialized")
	}
	// パスワードハッシュを比較
	if !u.passwordSvc.CheckPasswordHash(password, user.Password) {
		return "", errors.New("incorrect password")
	}
	// JWTを生成
	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userUsecase) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	// ユーザーをメールアドレスで取得
	return u.userRepo.GetUserByEmail(ctx, email)
}

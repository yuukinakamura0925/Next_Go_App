// internal/services/user_service.go
package services

import (
	"context"

	"Next_Go_App/ent"
	"Next_Go_App/internal/repositories"
)

// 現状使ってない。パスワードのハッシュ化を実装する際などに使う
// サービス層はリポジトリ層を利用して、ビジネスロジックを実装します。
type UserService struct {
	repo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignUp(ctx context.Context, email, name, password string) (*ent.User, error) {
	return s.repo.CreateUser(ctx, email, name, password)
}

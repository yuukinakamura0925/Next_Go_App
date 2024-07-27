// internal/repositories/user_repository.go
package repositories

import (
	"context"

	"Next_Go_App/ent"
	"Next_Go_App/ent/user"
)

// UserRepository はユーザー関連のデータ操作を定義するインターフェースです。
type IUserRepository interface {
	CreateUser(ctx context.Context, email, name, password string) (*ent.User, error)
	GetUserByEmailAndPassword(ctx context.Context, email, password string) (*ent.User, error)
}

// userRepository は UserRepository インターフェースを実装する構造体です。
type userRepository struct {
	client *ent.Client
}

// NewUserRepository は新しい userRepository のインスタンスを返します。
func NewUserRepository(client *ent.Client) IUserRepository {
	return &userRepository{client: client}
}

// CreateUser は新しいユーザーを作成します。
func (r *userRepository) CreateUser(ctx context.Context, email, name, password string) (*ent.User, error) {
	return r.client.User.
		Create().
		SetEmail(email).
		SetName(name).
		SetPassword(password).
		Save(ctx)
}

// GetUserByEmailAndPassword は email と password に一致するユーザーを取得します。
func (r *userRepository) GetUserByEmailAndPassword(ctx context.Context, email, password string) (*ent.User, error) {
	return r.client.User.
		Query().
		Where(user.EmailEQ(email), user.PasswordEQ(password)).
		First(ctx)
}

package repository

import (
	"Next_Go_App/domain/entities"
	"context"
)

// repository ディレクトリには、エンティティのデータ操作を行うインターフェースが含まれます。
type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	// 他の必要なメソッドもここに追加します（例: GetUserByEmail, UpdateUser, DeleteUserなど）
}

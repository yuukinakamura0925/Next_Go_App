package services

import "golang.org/x/crypto/bcrypt"

// PasswordService はパスワードのハッシュ化と検証のためのサービスを定義します。
type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type passwordService struct{}

// NewPasswordService は新しい PasswordService を返します。
func NewPasswordService() PasswordService {
	return &passwordService{}
}

// HashPassword はパスワードをハッシュ化します。
func (s *passwordService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash はパスワードとそのハッシュが一致するか検証します。
func (s *passwordService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

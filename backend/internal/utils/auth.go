package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// init関数はパッケージが初期化されたときに呼び出されます。
func init() {
	// .envファイルをロードします。
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// JWTの秘密鍵を環境変数から取得します。
var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Claims はJWTのペイロードを表します。
type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateJWT は指定されたメールアドレスを含むJWTトークンを生成します。
func GenerateJWT(id int, email string) (string, error) {
	// トークンの有効期限を設定します。
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: id,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// トークンを生成し、署名します。
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateJWT は指定されたJWTトークンを検証します。
func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}

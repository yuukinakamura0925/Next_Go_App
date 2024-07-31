package middlewares

import (
	"Next_Go_App/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware はJWTトークンを検証するミドルウェアです。
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエストヘッダーからAuthorizationヘッダーを取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// Authorizationヘッダーがない場合は401 Unauthorizedを返し、リクエストを中止
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			c.Abort()
			return
		}

		// Authorizationヘッダーの値を"Bearer "で分割してトークンを取得
		tokenStr := strings.Split(authHeader, "Bearer ")[1]
		// トークンを検証してクレーム(claims)を取得
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			// トークンが無効な場合は401 Unauthorizedを返し、リクエストを中止
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// クレームからemailをコンテキストに設定
		c.Set("email", claims.Email)
		// 次のミドルウェアまたはハンドラを実行
		c.Next()
	}
}

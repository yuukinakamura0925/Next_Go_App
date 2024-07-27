package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
)

// CORS設定を行い、他のオリジンからのリクエストを許可します。
func CORSMiddleware() gin.HandlerFunc {
	// 環境変数からCORS許可オリジンを取得
	corsAllowedOrigin := os.Getenv("CORS_ALLOWED_ORIGIN")
	if corsAllowedOrigin == "" {
		corsAllowedOrigin = "http://localhost:3000" // デフォルト値
	}

	// CORS設定
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", corsAllowedOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

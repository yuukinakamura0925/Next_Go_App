package main

import (
	"Next_Go_App/internal/db"
	"Next_Go_App/internal/handlers"
	"Next_Go_App/internal/middlewares"
	"Next_Go_App/internal/repositories"
	"Next_Go_App/internal/services"
	"Next_Go_App/internal/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginフレームワークのデフォルトの設定を使用してルータを作成
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	client := db.NewClient()

	// ユーザー関連の依存関係を設定
	userRepo := repositories.NewUserRepository(client)
	passwordService := services.NewPasswordService()
	userUsecase := usecases.NewUserUsecase(userRepo, passwordService)

	// 本関連の依存関係を設定
	bookRepo := repositories.NewBookRepository(client)
	bookUsecase := usecases.NewBookUsecase(bookRepo)

	// APIルートグループを定義
	api := router.Group("/api")
	// 認証不要の公開APIルート
	authApi := router.Group("/api/auth")
	// 認証が必要な保護されたAPIルート
	api.Use(middlewares.AuthMiddleware())

	// ハンドラを設定
	handlers.NewUserHandler(api, authApi, userUsecase)
	handlers.NewBookHandler(api, bookUsecase)

	// ルートハンドラの定義
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// サーバー起動
	router.Run(":8080")
}

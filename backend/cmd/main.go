package main

// 各レイヤーの関係性
// main.go:

// リポジトリ層 (repositories):
// データベースクライアントに依存し、データの保存と取得を行う。

// サービス層 (services):
// リポジトリ層に依存し、ビジネスロジックを提供する。

// ハンドラ層 (handlers):
// サービス層に依存し、HTTPリクエストとレスポンスを処理する。

// ミドルウェア (middlewares):
// 他のレイヤーに依存せず、リクエストとレスポンスに共通の処理を追加する。

import (
	// データベースとの接続を管理します。
	"Next_Go_App/internal/db"
	"Next_Go_App/internal/services"

	// HTTPリクエストとレスポンスの処理を行います。
	"Next_Go_App/internal/handlers"
	// ミドルウェアの定義を行います。
	"Next_Go_App/internal/middlewares"
	// データの保存と取得に関するロジックをカプセル化します。
	"Next_Go_App/internal/repositories"
	// シナリオロジックを含むユースケース層です。
	"Next_Go_App/internal/usecases"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	handlers.NewUserHandler(router, userUsecase)

	// 本関連の依存関係を設定
	bookRepo := repositories.NewBookRepository(client)
	bookUsecase := usecases.NewBookUsecase(bookRepo)
	handlers.NewBookHandler(router, bookUsecase)

	// 認証ミドルウェアを使用して保護されたAPIルートを作成できる
	// TODO: サインアップ、サインイン以外は認証が必要なルートに含めていくべき
	// auth := router.Group("/")
	// auth.Use(middlewares.AuthMiddleware())
	// auth.GET("/protected", func(c *gin.Context) {
	// 	email := c.GetString("email")
	// 	c.JSON(200, gin.H{"message": "protected route", "email": email})
	// })

	// ルートハンドラの定義
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// サーバー起動
	router.Run(":8080")
}

package main

// 各レイヤーの関係性
// main.go:

// db.NewClient()を呼び出してデータベース接続を設定。
// repositories.NewUserRepository(client)を呼び出してリポジトリ層を設定。
// services.NewUserService(userRepo)を呼び出してサービス層を設定。
// handlers.NewUserHandler(router, userService)を呼び出してハンドラ層を設定。
// router.Use(middlewares.CORSMiddleware())でミドルウェアを設定。

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

	// internal/middlewaresパッケージのCORSMiddleware関数を使用してCORSミドルウェアを設定
	// 1.リクエストはまずここを通る
	router.Use(middlewares.CORSMiddleware())

	// internal/dbパッケージのNewClient関数を使用してデータベースクライアントを作成
	client := db.NewClient()

	// internal/repositoriesパッケージのNewUserRepository関数を使用してUserRepositoryを作成
	userRepo := repositories.NewUserRepository(client) // 4.ユースケース層で呼ばれるリポジトリ層の対応するメソッドが呼び出されます。実際のDBへの接続がここで行われる
	// internal/usecasesパッケージのNewUserUsecase関数を使用してUserUsecaseを作成
	userUsecase := usecases.NewUserUsecase(userRepo) // 3.2でリクエストに対してハンドラが呼び出されると、その中でユースケース層の対応するメソッドが呼び出されます。
	// internal/handlersパッケージのNewUserHandler関数を使用してUserHandlerを作成
	handlers.NewUserHandler(router, userUsecase) // 2.リクエストは internal/handlers に定義された各ハンドラに渡されます。

	bookRepo := repositories.NewBookRepository(client)
	bookUsecase := usecases.NewBookUsecase(bookRepo)
	handlers.NewBookHandler(router, bookUsecase)

	// ルートハンドラの定義
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// サーバー起動
	router.Run(":8080")
}

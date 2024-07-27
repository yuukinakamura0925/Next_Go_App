package db

import (
	"context"
	"log"

	"Next_Go_App/ent"

	_ "github.com/lib/pq"
)

// PostgreSQLデータベースへの接続を確立し、スキーマを自動的にマイグレーションします。
func NewClient() *ent.Client {
	// PostgreSQLに接続
	client, err := ent.Open("postgres", "host=db port=5432 user=postgres dbname=db password=password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// データベーススキーマの自動マイグレーション
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

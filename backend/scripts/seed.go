package main

import (
	"Next_Go_App/ent"
	"Next_Go_App/internal/db"
	"context"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	client := db.NewClient()
	defer client.Close()

	// クライアントがnilでないことを確認
	if client == nil {
		log.Fatal("Database client is nil")
	}

	if err := SeedData(client); err != nil {
		log.Fatalf("failed seeding data: %v", err)
	}
}

// SeedData seeds the database with initial data
func SeedData(client *ent.Client) error {
	ctx := context.Background()

	// Debug: clientがnilでないことを確認
	if client == nil {
		log.Fatal("Ent client is nil")
	}

	// Sample user creation
	userCreate := client.User.Create()
	if userCreate == nil {
		log.Fatal("UserCreate is nil")
	}

	_, err := userCreate.
		SetName("John Doe").
		SetEmail("john.doe@example.com").
		SetPassword("secret").
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

package testutil

import (
	"context"
	"testing"

	"zach-sikora-daycare/internal/config"
	"zach-sikora-daycare/internal/database"
)

func NewTestDB(t *testing.T) *database.DB {
	t.Helper()

	ctx := context.Background()
	db, err := database.New(ctx, ":memory:")
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	return db
}

func NewTestConfig(t *testing.T) *config.Config {
	t.Helper()

	return &config.Config{
		DatabaseURL: ":memory:",
		Port:        "0",
		Env:         "test",
		Site: config.SiteConfig{
			Name: "Kids First Childcare",
			URL:  "http://localhost:3000",
		},
	}
}

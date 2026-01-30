package database

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"

	"zach-sikora-daycare/internal/database/sqlc"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

var gooseSetup sync.Once

type DB struct {
	Conn    *sql.DB
	Queries *sqlc.Queries
}

func New(ctx context.Context, databasePath string) (*DB, error) {
	if databasePath != ":memory:" {
		dir := filepath.Dir(databasePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("unable to create database directory: %w", err)
		}
	}

	dsn := databasePath
	if databasePath != ":memory:" {
		dsn = databasePath + "?_foreign_keys=on&_journal_mode=WAL"
	}

	conn, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %w", err)
	}

	if err := conn.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	db := &DB{
		Conn:    conn,
		Queries: sqlc.New(conn),
	}

	if err := db.migrate(); err != nil {
		return nil, fmt.Errorf("unable to run migrations: %w", err)
	}

	return db, nil
}

func (db *DB) Close() {
	db.Conn.Close()
}

func (db *DB) migrate() error {
	gooseSetup.Do(func() {
		goose.SetBaseFS(migrationsFS)
		_ = goose.SetDialect("sqlite3")
	})

	if err := goose.Up(db.Conn, "migrations"); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

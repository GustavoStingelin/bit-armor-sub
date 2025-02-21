package internal

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"os"
)

var DBPool *pgxpool.Pool

func InitDBPool() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, databaseName)

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		slog.Error("Failed to parse database URL: %v", err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	DBPool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		slog.Error("Failed to create connection pool: %v", err)
	}

	// test connection
	var version string
	err = DBPool.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		slog.Error("Query failed", "err", err)
	}

	slog.Info("Connected to pg", "version", version)
}

func CloseDBPool() {
	if DBPool != nil {
		DBPool.Close()
	}
}

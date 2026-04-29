package database

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-ecommerce-pkg/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// RunMigrations executes database migrations using goose with the default "DB" prefix.
// path: directory containing migration files.
func RunMigrations(log logger.LoggerInterface, path string) error {
	return RunMigrationsWithPrefix(log, path, "DB")
}

// RunMigrationsWithPrefix executes database migrations using goose against a
// prefix-specific database configuration. This enables per-service database
// migrations by reading DB_<PREFIX>_HOST, DB_<PREFIX>_PORT, etc. from env.
func RunMigrationsWithPrefix(log logger.LoggerInterface, path string, prefix string) error {
	if prefix == "" {
		prefix = "DB"
	}

	host := viper.GetString(fmt.Sprintf("%s_HOST", prefix))
	if host == "" {
		host = viper.GetString("DB_HOST")
	}
	port := viper.GetString(fmt.Sprintf("%s_PORT", prefix))
	if port == "" {
		port = viper.GetString("DB_PORT")
	}
	user := viper.GetString(fmt.Sprintf("%s_USERNAME", prefix))
	if user == "" {
		user = viper.GetString("DB_USERNAME")
	}
	dbname := viper.GetString(fmt.Sprintf("%s_NAME", prefix))
	if dbname == "" {
		dbname = viper.GetString("DB_NAME")
	}
	password := viper.GetString(fmt.Sprintf("%s_PASSWORD", prefix))
	if password == "" {
		password = viper.GetString("DB_PASSWORD")
	}

	// Use pgx driver for goose
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password,
	)

	db, err := goose.OpenDBWithDriver("pgx", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database for migrations (prefix=%s): %w", prefix, err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Error("Failed to close database after migrations", zap.Error(err), zap.String("prefix", prefix))
		}
	}()

	log.Info("Running database migrations",
		zap.String("path", path),
		zap.String("dbname", dbname),
		zap.String("prefix", prefix),
	)

	if err := goose.RunContext(context.Background(), "up", db, path); err != nil {
		return fmt.Errorf("migration 'up' failed (prefix=%s): %w", prefix, err)
	}

	log.Info("Database migrations completed successfully",
		zap.String("prefix", prefix),
		zap.String("dbname", dbname),
	)
	return nil
}

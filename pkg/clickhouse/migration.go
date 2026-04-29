package clickhouse

import (
	"context"
	"fmt"

	_ "github.com/ClickHouse/clickhouse-go/v2" // Import database/sql driver
	"github.com/MamangRust/monolith-ecommerce-pkg/logger"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// RunMigrations executes ClickHouse schema migrations using goose.
func RunMigrations(log logger.LoggerInterface, path string) error {
	addr := viper.GetString("CLICKHOUSE_ADDR")
	if addr == "" {
		host := viper.GetString("CLICKHOUSE_HOST")
		if host == "" {
			host = "clickhouse"
		}
		port := viper.GetString("CLICKHOUSE_PORT")
		if port == "" {
			port = "9000"
		}
		addr = fmt.Sprintf("%s:%s", host, port)
	}

	dbName := viper.GetString("CLICKHOUSE_DATABASE")
	user := viper.GetString("CLICKHOUSE_USERNAME")
	password := viper.GetString("CLICKHOUSE_PASSWORD")

	connStr := fmt.Sprintf("clickhouse://%s:%s@%s/%s?dial_timeout=30s&max_execution_time=60", 
		user, password, addr, dbName)

	if err := goose.SetDialect("clickhouse"); err != nil {
		return fmt.Errorf("failed to set goose clickhouse dialect: %w", err)
	}

	db, err := goose.OpenDBWithDriver("clickhouse", connStr)
	if err != nil {
		return fmt.Errorf("failed to open clickhouse database for migrations: %w", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Error("Failed to close clickhouse database after migrations", zap.Error(err))
		}
	}()

	log.Info("Running clickhouse migrations",
		zap.String("path", path),
		zap.String("dbname", dbName),
	)

	if err := goose.RunContext(context.Background(), "up", db, path); err != nil {
		return fmt.Errorf("clickhouse migration 'up' failed: %w", err)
	}

	log.Info("ClickHouse database migrations completed successfully")
	return nil
}

package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/MamangRust/microservice-ecommerce-pkg/clickhouse"
	"github.com/MamangRust/microservice-ecommerce-pkg/dotenv"
	"github.com/MamangRust/microservice-ecommerce-pkg/kafka"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-grpc/service/stats-writer/handler"
	"github.com/MamangRust/microservice-ecommerce-grpc/service/stats-writer/repository"
	"github.com/MamangRust/microservice-ecommerce-grpc/service/stats-writer/usecase"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	if err := dotenv.Viper(); err != nil {
		zap.L().Error("Failed to load configuration", zap.Error(err))
	}
	log, _ := logger.NewLogger("stats-writer", nil)

	// Run ClickHouse Migrations
	migrationPath := "./database/migrations"
	if err := clickhouse.RunMigrations(log, migrationPath); err != nil {
		log.Fatal("Failed to run ClickHouse migrations", zap.Error(err))
	}

	chConn, err := clickhouse.NewClient(log)
	if err != nil {
		log.Fatal("Failed to connect to ClickHouse", zap.Error(err))
	}

	brokers := strings.Split(viper.GetString("KAFKA_BROKERS"), ",")
	if len(brokers) == 0 || brokers[0] == "" {
		brokers = []string{"kafka-1:9092", "kafka-2:9092", "kafka-3:9092"}
	}
	k := kafka.NewKafka(log, brokers)

	// Dependency Injection
	repo := repository.NewClickhouseRepository(chConn, log)
	uc := usecase.NewStatsUseCase(repo)
	statsHandler := handler.NewStatsHandler(uc, log)

	// Start Consumer
	// Ecommerce topics: category_stats_topic, order_stats_topic, transaction_stats_topic
	topics := []string{
		"category_stats_topic",
		"order_stats_topic",
		"transaction_stats_topic",
	}
	
	if err := k.StartConsumers(topics, "stats-writer-group", statsHandler); err != nil {
		log.Fatal("Failed to start Kafka consumers", zap.Error(err))
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down Stats Writer...")
}

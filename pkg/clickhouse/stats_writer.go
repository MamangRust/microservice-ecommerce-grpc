package clickhouse

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/IBM/sarama"
	"github.com/MamangRust/microservice-ecommerce-pkg/kafka"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"go.uber.org/zap"
)

// Events
type CategoryStatEvent struct {
	CategoryID        uint32 `json:"category_id"`
	TotalViews        uint32 `json:"total_views"`
	TotalTransactions uint32 `json:"total_transactions"`
}

type OrderStatEvent struct {
	OrderID      uint32 `json:"order_id"`
	TotalRevenue uint64 `json:"total_revenue"`
	ItemsCount   uint32 `json:"items_count"`
}

type TransactionStatEvent struct {
	TransactionID uint32 `json:"transaction_id"`
	PaymentMethod string `json:"payment_method"`
	Amount        uint64 `json:"amount"`
	Status        string `json:"status"`
}

type StatsWriter struct {
	conn   clickhouse.Conn
	logger logger.LoggerInterface
	kafka  *kafka.Kafka
}

func NewStatsWriter(conn clickhouse.Conn, l logger.LoggerInterface, k *kafka.Kafka) *StatsWriter {
	return &StatsWriter{
		conn:   conn,
		logger: l,
		kafka:  k,
	}
}

func (w *StatsWriter) Start() error {
	topics := []string{"category_stats_topic", "order_stats_topic", "transaction_stats_topic"}
	
	w.logger.Info("Starting ClickHouse StatsWriter consumer", zap.Strings("topics", topics))
	
	return w.kafka.StartConsumers(topics, "clickhouse_stats_writer_group", w)
}

func (w *StatsWriter) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (w *StatsWriter) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (w *StatsWriter) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := context.Background()
	for message := range claim.Messages() {
		w.logger.Debug("Received stats event", zap.String("topic", message.Topic))

		var err error
		switch message.Topic {
		case "category_stats_topic":
			err = w.handleCategoryStat(ctx, message.Value)
		case "order_stats_topic":
			err = w.handleOrderStat(ctx, message.Value)
		case "transaction_stats_topic":
			err = w.handleTransactionStat(ctx, message.Value)
		}

		if err != nil {
			w.logger.Error("Failed to handle stats event", zap.Error(err), zap.String("topic", message.Topic))
		} else {
			session.MarkMessage(message, "")
		}
	}
	return nil
}

func (w *StatsWriter) handleCategoryStat(ctx context.Context, data []byte) error {
	var event CategoryStatEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	query := `INSERT INTO category_stats (date, category_id, total_views, total_transactions) VALUES (?, ?, ?, ?)`
	return w.conn.Exec(ctx, query, time.Now(), event.CategoryID, event.TotalViews, event.TotalTransactions)
}

func (w *StatsWriter) handleOrderStat(ctx context.Context, data []byte) error {
	var event OrderStatEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	query := `INSERT INTO order_stats (date, order_id, total_revenue, items_count) VALUES (?, ?, ?, ?)`
	return w.conn.Exec(ctx, query, time.Now(), event.OrderID, event.TotalRevenue, event.ItemsCount)
}

func (w *StatsWriter) handleTransactionStat(ctx context.Context, data []byte) error {
	var event TransactionStatEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	query := `INSERT INTO transaction_stats (date, transaction_id, payment_method, amount, status) VALUES (?, ?, ?, ?, ?)`
	return w.conn.Exec(ctx, query, time.Now(), event.TransactionID, event.PaymentMethod, event.Amount, event.Status)
}

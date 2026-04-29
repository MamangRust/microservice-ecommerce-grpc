package clickhouse

import (
	"context"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"go.uber.org/zap"
)

type StatsReader interface {
	GetCategoryStats(ctx context.Context, categoryID uint32, start, end time.Time) (views uint32, transactions uint32, err error)
	GetOrderStats(ctx context.Context, orderID uint32, start, end time.Time) (revenue uint64, items uint32, err error)
	GetTransactionStats(ctx context.Context, transactionID uint32, start, end time.Time) (amount uint64, count uint32, err error)
}

type statsReaderImpl struct {
	conn   clickhouse.Conn
	logger logger.LoggerInterface
}

func NewStatsReader(conn clickhouse.Conn, l logger.LoggerInterface) StatsReader {
	return &statsReaderImpl{
		conn:   conn,
		logger: l,
	}
}

func (r *statsReaderImpl) GetCategoryStats(ctx context.Context, categoryID uint32, start, end time.Time) (uint32, uint32, error) {
	query := `
		SELECT sum(total_views), sum(total_transactions)
		FROM category_stats
		WHERE category_id = ? AND date >= ? AND date <= ?
	`
	row := r.conn.QueryRow(ctx, query, categoryID, start.Format("2006-01-02"), end.Format("2006-01-02"))
	
	var views, txs uint64
	if err := row.Scan(&views, &txs); err != nil {
		r.logger.Error("Failed to query category stats", zap.Error(err))
		// ClickHouse returns 0 for SUM on empty sets, but could still error on closed connections
		return 0, 0, err
	}
	return uint32(views), uint32(txs), nil
}

func (r *statsReaderImpl) GetOrderStats(ctx context.Context, orderID uint32, start, end time.Time) (uint64, uint32, error) {
	query := `
		SELECT sum(total_revenue), sum(items_count)
		FROM order_stats
		WHERE order_id = ? AND date >= ? AND date <= ?
	`
	row := r.conn.QueryRow(ctx, query, orderID, start.Format("2006-01-02"), end.Format("2006-01-02"))
	
	var revenue uint64
	var items uint64
	if err := row.Scan(&revenue, &items); err != nil {
		r.logger.Error("Failed to query order stats", zap.Error(err))
		return 0, 0, err
	}
	return revenue, uint32(items), nil
}

func (r *statsReaderImpl) GetTransactionStats(ctx context.Context, transactionID uint32, start, end time.Time) (uint64, uint32, error) {
	query := `
		SELECT sum(amount), count(*)
		FROM transaction_stats
		WHERE transaction_id = ? AND date >= ? AND date <= ?
	`
	row := r.conn.QueryRow(ctx, query, transactionID, start.Format("2006-01-02"), end.Format("2006-01-02"))
	
	var amount uint64
	var count uint64
	if err := row.Scan(&amount, &count); err != nil {
		r.logger.Error("Failed to query transaction stats", zap.Error(err))
		return 0, 0, err
	}
	return amount, uint32(count), nil
}

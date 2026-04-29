package repository

import (
	"context"
	"time"

	chDriver "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/events"
)

type clickhouseRepository struct {
	conn chDriver.Conn
	log  logger.LoggerInterface
}

func NewClickhouseRepository(conn chDriver.Conn, log logger.LoggerInterface) Repository {
	return &clickhouseRepository{
		conn: conn,
		log:  log,
	}
}

func (r *clickhouseRepository) InsertCategoryStat(ctx context.Context, event events.CategoryStatEvent) error {
	query := `INSERT INTO category_stats (date, category_id, category_name, merchant_id, total_views, total_transactions, total_price, items_sold, order_count) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	return r.conn.Exec(ctx, query, time.Now(), event.CategoryID, event.CategoryName, event.MerchantID, event.TotalViews, event.TotalTransactions, event.TotalPrice, event.ItemsSold, event.OrderCount)
}

func (r *clickhouseRepository) InsertOrderStat(ctx context.Context, event events.OrderStatEvent) error {
	query := `INSERT INTO order_stats (date, order_id, merchant_id, total_revenue, total_items_sold, active_cashiers, unique_products_sold) VALUES (?, ?, ?, ?, ?, ?, ?)`
	return r.conn.Exec(ctx, query, time.Now(), event.OrderID, event.MerchantID, event.TotalRevenue, event.TotalItemsSold, event.ActiveCashiers, event.UniqueProductsSold)
}

func (r *clickhouseRepository) InsertTransactionStat(ctx context.Context, event events.TransactionStatEvent) error {
	query := `INSERT INTO transaction_stats (date, transaction_id, merchant_id, payment_method, amount, status, total_amount) VALUES (?, ?, ?, ?, ?, ?, ?)`
	return r.conn.Exec(ctx, query, time.Now(), event.TransactionID, event.MerchantID, event.PaymentMethod, event.Amount, event.Status, event.TotalAmount)
}

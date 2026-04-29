package events

import "time"

const (
	CategoryStatsTopic    = "category_stats_topic"
	OrderStatsTopic       = "order_stats_topic"
	TransactionStatsTopic = "transaction_stats_topic"
)

type CategoryStatEvent struct {
	CategoryID        uint32    `json:"category_id"`
	CategoryName      string    `json:"category_name"`
	MerchantID        uint32    `json:"merchant_id"`
	TotalViews        uint32    `json:"total_views"`
	TotalTransactions uint32    `json:"total_transactions"`
	TotalPrice        uint64    `json:"total_price"`
	ItemsSold         uint32    `json:"items_sold"`
	OrderCount        uint32    `json:"order_count"`
	CreatedAt         time.Time `json:"created_at"`
}

type OrderStatEvent struct {
	OrderID            uint32    `json:"order_id"`
	MerchantID         uint32    `json:"merchant_id"`
	TotalRevenue       uint64    `json:"total_revenue"`
	TotalItemsSold     uint32    `json:"total_items_sold"`
	ActiveCashiers     uint32    `json:"active_cashiers"`
	UniqueProductsSold uint32    `json:"unique_products_sold"`
	CreatedAt          time.Time `json:"created_at"`
}

type TransactionStatEvent struct {
	TransactionID uint32    `json:"transaction_id"`
	OrderID       uint32    `json:"order_id"`
	UserID        uint32    `json:"user_id"`
	MerchantID    uint32    `json:"merchant_id"`
	PaymentMethod string    `json:"payment_method"`
	Amount        uint64    `json:"amount"`
	TotalAmount   uint64    `json:"total_amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

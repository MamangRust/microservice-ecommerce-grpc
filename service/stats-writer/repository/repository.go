package repository

import (
	"context"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/events"
)

type Repository interface {
	InsertCategoryStat(ctx context.Context, event events.CategoryStatEvent) error
	InsertOrderStat(ctx context.Context, event events.OrderStatEvent) error
	InsertTransactionStat(ctx context.Context, event events.TransactionStatEvent) error
}

package service

import (
	"context"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
)

type OrderQueryService interface {
	FindAll(
		ctx context.Context,
		req *requests.FindAllOrder,
	) ([]*db.GetOrdersRow, *int, error)

	FindActive(
		ctx context.Context,
		req *requests.FindAllOrder,
	) ([]*db.GetOrdersActiveRow, *int, error)

	FindTrashed(
		ctx context.Context,
		req *requests.FindAllOrder,
	) ([]*db.GetOrdersTrashedRow, *int, error)

	FindByID(
		ctx context.Context,
		order_id int,
	) (*db.GetOrderByIDRow, error)
}

type OrderCommandService interface {
	Create(
		ctx context.Context,
		request *requests.CreateOrderRequest,
	) (*db.CreateOrderRow, error)

	Update(
		ctx context.Context,
		request *requests.UpdateOrderRequest,
	) (*db.UpdateOrderRow, error)

	Trash(
		ctx context.Context,
		order_id int,
	) (*db.Order, error)

	Restore(
		ctx context.Context,
		order_id int,
	) (*db.Order, error)

	DeletePermanent(
		ctx context.Context,
		order_id int,
	) (bool, error)
	RestoreAll(ctx context.Context) (bool, error)
	DeleteAll(ctx context.Context) (bool, error)
}

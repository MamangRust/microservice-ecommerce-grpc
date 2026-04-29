package repository

import (
	"context"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
)

type CategoryQueryRepository interface {
	FindAll(ctx context.Context, req *requests.FindAllCategory) ([]*db.GetCategoriesRow, error)

	FindActive(ctx context.Context, req *requests.FindAllCategory) ([]*db.GetCategoriesActiveRow, error)

	FindTrashed(ctx context.Context, req *requests.FindAllCategory) ([]*db.GetCategoriesTrashedRow, error)

	FindByID(ctx context.Context, category_id int) (*db.GetCategoryByIDRow, error)
	FindByIDTrashed(ctx context.Context, category_id int) (*db.Category, error)
}

type CategoryCommandRepository interface {
	Create(
		ctx context.Context,
		request *requests.CreateCategoryRequest,
	) (*db.CreateCategoryRow, error)

	Update(
		ctx context.Context,
		request *requests.UpdateCategoryRequest,
	) (*db.UpdateCategoryRow, error)

	Trash(
		ctx context.Context,
		category_id int,
	) (*db.Category, error)

	Restore(
		ctx context.Context,
		category_id int,
	) (*db.Category, error)

	DeletePermanent(
		ctx context.Context,
		category_id int,
	) (bool, error)

	RestoreAll(ctx context.Context) (bool, error)
	DeleteAll(ctx context.Context) (bool, error)
}

package usecase

import (
	"context"
	"github.com/MamangRust/microservice-ecommerce-grpc/service/stats-writer/repository"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/events"
)

type UseCase interface {
	SaveCategoryStat(ctx context.Context, event events.CategoryStatEvent) error
	SaveOrderStat(ctx context.Context, event events.OrderStatEvent) error
	SaveTransactionStat(ctx context.Context, event events.TransactionStatEvent) error
}

type statsUseCase struct {
	repo repository.Repository
}

func NewStatsUseCase(repo repository.Repository) UseCase {
	return &statsUseCase{
		repo: repo,
	}
}

func (u *statsUseCase) SaveCategoryStat(ctx context.Context, event events.CategoryStatEvent) error {
	return u.repo.InsertCategoryStat(ctx, event)
}

func (u *statsUseCase) SaveOrderStat(ctx context.Context, event events.OrderStatEvent) error {
	return u.repo.InsertOrderStat(ctx, event)
}

func (u *statsUseCase) SaveTransactionStat(ctx context.Context, event events.TransactionStatEvent) error {
	return u.repo.InsertTransactionStat(ctx, event)
}

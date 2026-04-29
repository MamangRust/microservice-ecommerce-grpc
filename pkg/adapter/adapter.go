package adapter

import (
	"context"

	db "github.com/MamangRust/microservice-ecommerce-pkg/database/schema"
	user_repo "github.com/MamangRust/microservice-ecommerce-grpc-user/repository"
)

type UserAdapter interface {
	GetByEmail(ctx context.Context, email string) (*db.User, error)
}

type localUserAdapter struct {
	repo user_repo.UserQueryRepository
}

func NewLocalUserAdapter(repo user_repo.UserQueryRepository) UserAdapter {
	return &localUserAdapter{repo: repo}
}

func (a *localUserAdapter) GetByEmail(ctx context.Context, email string) (*db.User, error) {
	return a.repo.FindByEmail(ctx, email)
}

// CardAdapter
type CardAdapter interface{}

type localCardAdapter struct{}

func NewLocalCardAdapter(queryRepo interface{}, commandRepo interface{}) CardAdapter {
	return &localCardAdapter{}
}

// MerchantAdapter
type MerchantAdapter interface{}

type localMerchantAdapter struct{}

func NewLocalMerchantAdapter(repo interface{}) MerchantAdapter {
	return &localMerchantAdapter{}
}

// SaldoAdapter
type SaldoAdapter interface{}

type localSaldoAdapter struct{}

func NewLocalSaldoAdapter(repos interface{}) SaldoAdapter {
	return &localSaldoAdapter{}
}

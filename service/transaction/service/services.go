package service

import (
	mencache "github.com/MamangRust/microservice-ecommerce-grpc-transaction/cache"
	"github.com/MamangRust/microservice-ecommerce-grpc-transaction/repository"
	"github.com/MamangRust/microservice-ecommerce-pkg/kafka"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/observability"
)

type Service struct {
	TransactionQuery   TransactionQueryService
	TransactionCommand TransactionCommandService
}

type Deps struct {
	Kafka         *kafka.Kafka
	Cache         *mencache.Mencache
	Repositories  *repository.Repositories
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
}

func NewService(deps *Deps) *Service {
	return &Service{
		TransactionQuery: NewTransactionQueryService(&TransactionQueryServiceDeps{
			Observability: deps.Observability,
			Cache:         deps.Cache.TransactionQueryCache,
			Repository:    deps.Repositories.TransactionQuery,
			Logger:        deps.Logger,
		}),
		TransactionCommand: NewTransactionCommandService(&TransactionCommandServiceDeps{
			Observability:      deps.Observability,
			Kafka:              deps.Kafka,
			Cache:              deps.Cache.TransactionCommandCache,
			TransactionQuery:   deps.Repositories.TransactionQuery,
			TransactionCommand: deps.Repositories.TransactionCommand,
			UserQuery:          deps.Repositories.UserQuery,
			MerchantQuery:      deps.Repositories.MerchantQuery,
			OrderQuery:         deps.Repositories.OrderQuery,
			OrderItem:          deps.Repositories.OrderItem,
			ShippingAddress:    deps.Repositories.ShippingAddress,
			Logger:             deps.Logger,
		}),
	}
}

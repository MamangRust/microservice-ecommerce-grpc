package service

import (
	mencache "github.com/MamangRust/microservice-ecommerce-grpc-order/cache"
	"github.com/MamangRust/microservice-ecommerce-grpc-order/repository"
	"github.com/MamangRust/microservice-ecommerce-pkg/kafka"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/observability"
)

type Service struct {
	OrderQuery   OrderQueryService
	OrderCommand OrderCommandService
}

type Deps struct {
	Cache         *mencache.Mencache
	Repositories  *repository.Repositories
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
	Kafka         *kafka.Kafka
}

func NewService(deps *Deps) *Service {
	return &Service{
		OrderQuery: NewOrderQueryService(&OrderQueryServiceDeps{
			Observability:   deps.Observability,
			Cache:           deps.Cache.OrderQueryCache,
			OrderRepository: deps.Repositories.OrderQuery,
			Logger:          deps.Logger,
		}),
		OrderCommand: NewOrderCommandService(&OrderCommandServiceDeps{
			Observability:              deps.Observability,
			Cache:                      deps.Cache.OrderCommandCache,
			UserQueryRepository:        deps.Repositories.UserQuery,
			ProductQueryRepository:     deps.Repositories.ProductQuery,
			ProductCommandRepository:   deps.Repositories.ProductCommand,
			OrderQueryRepository:       deps.Repositories.OrderQuery,
			OrderCommandRepository:     deps.Repositories.OrderCommand,
			OrderItemQueryRepository:   deps.Repositories.OrderItemQuery,
			OrderItemCommandRepository: deps.Repositories.OrderItemCommand,
			MerchantQueryRepository:    deps.Repositories.MerchantQuery,
			ShippingAddressRepository:  deps.Repositories.ShippingAddress,
			TransactionCommandRepository: deps.Repositories.TransactionCommand,
			ShippingQueryRepository:    deps.Repositories.ShippingQuery,
			Logger:                     deps.Logger,
			Kafka:                      deps.Kafka,
		}),
	}
}

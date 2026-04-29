package service

import (
	mencache "github.com/MamangRust/microservice-ecommerce-grpc-category/cache"
	"github.com/MamangRust/microservice-ecommerce-grpc-category/repository"
	"github.com/MamangRust/microservice-ecommerce-pkg/kafka"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/observability"
)

type Service struct {
	CategoryQuery   CategoryQueryService
	CategoryCommand CategoryCommandService
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
		CategoryQuery: NewCategoryQueryService(&CategoryQueryServiceDeps{
			Observability:           deps.Observability,
			Cache:                   deps.Cache.CategoryQueryCache,
			CategoryQueryRepository: deps.Repositories.CategoryQuery,
			Logger:                  deps.Logger,
		}),
		CategoryCommand: NewCategoryCommandService(&CategoryCommandServiceDeps{
			Observability: deps.Observability,
			Cache:         deps.Cache.CategoryCommandCache,
			CategoryCommandRepository: deps.Repositories.
				CategoryCommand,
			CategoryQueryRepository: deps.Repositories.CategoryQuery,
			Logger:                  deps.Logger,
			Kafka:                   deps.Kafka,
		}),
	}
}

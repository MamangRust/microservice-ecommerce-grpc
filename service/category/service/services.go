package service

import (
	mencache "github.com/MamangRust/monolith-ecommerce-grpc-category/cache"
	"github.com/MamangRust/monolith-ecommerce-grpc-category/repository"
	"github.com/MamangRust/monolith-ecommerce-pkg/kafka"
	"github.com/MamangRust/monolith-ecommerce-pkg/logger"
	"github.com/MamangRust/monolith-ecommerce-shared/observability"
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

package apps

import (
	"github.com/MamangRust/monolith-ecommerce-grpc-category/cache"
	"github.com/MamangRust/monolith-ecommerce-grpc-category/handler"
	"github.com/MamangRust/monolith-ecommerce-grpc-category/repository"
	"github.com/MamangRust/monolith-ecommerce-grpc-category/service"
	"github.com/MamangRust/monolith-ecommerce-pkg/kafka"
	"github.com/MamangRust/monolith-ecommerce-pkg/server"
	"github.com/MamangRust/monolith-ecommerce-shared/observability"
	"github.com/MamangRust/monolith-ecommerce-shared/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func NewServer(cfg *server.Config) (*server.GRPCServer, error) {
	srv, err := server.New(cfg)
	if err != nil {
		return nil, err
	}

	repos := repository.NewRepositories(srv.DB)
	obs, _ := observability.NewObservability("category-server", srv.Logger)
	myKafka := kafka.NewKafka(srv.Logger, []string{viper.GetString("KAFKA_BROKERS")})
	cache := cache.NewMencache(srv.CacheStore)

	svc := service.NewService(&service.Deps{
		Cache:         cache,
		Logger:        srv.Logger,
		Repositories:  repos,
		Observability: obs,
		Kafka:         myKafka,
	})

	h := handler.NewHandler(&handler.Deps{Service: svc, Logger: srv.Logger})

	srv.RegisterServices = func(gs *grpc.Server) {
		pb.RegisterCategoryQueryServiceServer(gs, h.CategoryQuery)
		pb.RegisterCategoryCommandServiceServer(gs, h.CategoryCommand)
	}

	return srv, nil
}

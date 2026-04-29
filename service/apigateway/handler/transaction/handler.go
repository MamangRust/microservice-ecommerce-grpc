package transactionhandler

import (
	transaction_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/cache/transaction"
	pb "github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	apimapper "github.com/MamangRust/microservice-ecommerce-shared/mapper/transaction"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	sharedErrors "github.com/MamangRust/microservice-ecommerce-shared/errors"
)

type DepsTransaction struct {
	Client      *grpc.ClientConn
	StatsReader *grpc.ClientConn
	E           *echo.Echo
	Logger      logger.LoggerInterface
	CacheStore  *cache.CacheStore
	ApiHandler  sharedErrors.ApiHandler
}

func RegisterTransactionHandler(deps *DepsTransaction) {
	mapper := apimapper.NewTransactionResponseMapper()
	statsMapper := apimapper.NewTransactionStatsResponseMapper()
	cache := transaction_cache.NewTransactionMencache(deps.CacheStore)

	queryClient := pb.NewTransactionQueryServiceClient(deps.Client)
	commandClient := pb.NewTransactionCommandServiceClient(deps.Client)
	statsClient := pb.NewTransactionStatsServiceClient(deps.StatsReader)
	statsByMerchantClient := pb.NewTransactionStatsByMerchantServiceClient(deps.StatsReader)

	NewTransactionQueryHandleApi(&transactionQueryHandleDeps{
		queryClient: queryClient,
		router:      deps.E,
		logger:      deps.Logger,
		mapper:      mapper.QueryMapper(),
		cache:       cache,
		apiHandler:  deps.ApiHandler,
	})

	NewTransactionCommandHandleApi(&transactionCommandHandleDeps{
		client:     commandClient,
		router:     deps.E,
		logger:     deps.Logger,
		mapper:     mapper.CommandMapper(),
		cache:      cache,
		apiHandler: deps.ApiHandler,
	})

	NewTransactionStatsHandleApi(&TransactionStatsHandleDeps{
		StatsClient:           statsClient,
		StatsByMerchantClient: statsByMerchantClient,
		Router:                deps.E,
		Logger:                deps.Logger,
		StatsMapper:           statsMapper,
		StatsCache:            cache,
		StatsByMerchantCache:  cache,
		ApiHandler:            deps.ApiHandler,
	})
}

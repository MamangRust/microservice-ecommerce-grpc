package reviewdetailhandler

import (
	reviewdetail_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/cache/review_detail"
	pb "github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-pkg/upload_image"
	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	apimapper "github.com/MamangRust/microservice-ecommerce-shared/mapper/review_detail"
	reviewapimapper "github.com/MamangRust/microservice-ecommerce-shared/mapper/review"
	"github.com/MamangRust/microservice-ecommerce-shared/observability"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type DepsReviewDetail struct {
	Client        *grpc.ClientConn
	E             *echo.Echo
	Logger        logger.LoggerInterface
	Cache         *cache.CacheStore
	Upload        upload_image.ImageUploads
	Observability observability.TraceLoggerObservability
}

func RegisterReviewDetailHandler(deps *DepsReviewDetail) {
	mapper := apimapper.NewReviewDetailResponseMapper()
	reviewMapper := reviewapimapper.NewReviewResponseMapper()
	cache := reviewdetail_cache.NewReviewDetailMencache(deps.Cache)

	NewReviewDetailQueryHandleApi(&reviewDetailQueryHandleDeps{
		client:        pb.NewReviewDetailQueryServiceClient(deps.Client),
		router:        deps.E,
		logger:        deps.Logger,
		mapper:        mapper.QueryMapper(),
		cache:         cache.QueryCache(),
		observability: deps.Observability,
	})

	NewReviewDetailCommandHandleApi(&reviewDetailCommandHandleDeps{
		client:        pb.NewReviewDetailCommandServiceClient(deps.Client),
		router:        deps.E,
		logger:        deps.Logger,
		mapper:        mapper.CommandMapper(),
		queryMapper:   mapper.QueryMapper(),
		reviewMapper:  reviewMapper.CommandMapper(),
		cache:         cache.CommandCache(),
		upload:        deps.Upload,
		observability: deps.Observability,
	})
}

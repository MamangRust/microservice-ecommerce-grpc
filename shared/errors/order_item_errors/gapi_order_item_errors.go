package orderitem_errors

import (
	"github.com/MamangRust/microservice-ecommerce-shared/errors"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcInvalidID = errors.NewGrpcError("invalid ID", int(codes.InvalidArgument))
)

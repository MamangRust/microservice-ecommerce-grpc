package handler

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-order/service"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	OrderQuery   OrderQueryHandler
	OrderCommand OrderCommandHandler
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		OrderQuery:   NewOrderQueryHandler(deps.Service.OrderQuery, deps.Logger),
		OrderCommand: NewOrderCommandHandler(deps.Service.OrderCommand, deps.Logger),
	}
}

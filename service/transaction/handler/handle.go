package handler

import (
	"github.com/MamangRust/monolith-ecommerce-grpc-transaction/service"
	"github.com/MamangRust/monolith-ecommerce-pkg/logger"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	TransactionQuery   TransactionQueryHandler
	TransactionCommand TransactionCommandHandler
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		TransactionQuery:   NewTransactionQueryHandler(deps.Service.TransactionQuery, deps.Logger),
		TransactionCommand: NewTransactionCommandHandler(deps.Service.TransactionCommand, deps.Logger),
	}
}

package handler

import "github.com/MamangRust/monolith-ecommerce-shared/pb"

type TransactionQueryHandler interface {
	pb.TransactionQueryServiceServer
}

type TransactionCommandHandler interface {
	pb.TransactionCommandServiceServer
}

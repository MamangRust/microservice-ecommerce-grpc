package handler

import (
	"context"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/microservice-ecommerce-grpc/service/stats-reader/repository"
)

type TransactionStatsHandler struct {
	pb.UnimplementedTransactionStatsServiceServer
	repo repository.Repository
	log  logger.LoggerInterface
}

func NewTransactionStatsHandler(repo repository.Repository, log logger.LoggerInterface) *TransactionStatsHandler {
	return &TransactionStatsHandler{
		repo: repo,
		log:  log,
	}
}

func (h *TransactionStatsHandler) GetMonthlyAmountSuccess(ctx context.Context, req *pb.MonthAmountTransactionRequest) (*pb.ApiResponseTransactionMonthAmountSuccess, error) {
	data, err := h.repo.GetMonthlyAmountSuccess(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionMonthAmountSuccess{
		Status:  "success",
		Message: "success get monthly amount success",
		Data:    data,
	}, nil
}

func (h *TransactionStatsHandler) GetYearlyAmountSuccess(ctx context.Context, req *pb.YearAmountTransactionRequest) (*pb.ApiResponseTransactionYearAmountSuccess, error) {
	data, err := h.repo.GetYearlyAmountSuccess(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionYearAmountSuccess{
		Status:  "success",
		Message: "success get yearly amount success",
		Data:    data,
	}, nil
}

func (h *TransactionStatsHandler) GetMonthlyAmountFailed(ctx context.Context, req *pb.MonthAmountTransactionRequest) (*pb.ApiResponseTransactionMonthAmountFailed, error) {
	data, err := h.repo.GetMonthlyAmountFailed(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionMonthAmountFailed{
		Status:  "success",
		Message: "success get monthly amount failed",
		Data:    data,
	}, nil
}

func (h *TransactionStatsHandler) GetYearlyAmountFailed(ctx context.Context, req *pb.YearAmountTransactionRequest) (*pb.ApiResponseTransactionYearAmountFailed, error) {
	data, err := h.repo.GetYearlyAmountFailed(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionYearAmountFailed{
		Status:  "success",
		Message: "success get yearly amount failed",
		Data:    data,
	}, nil
}

func (h *TransactionStatsHandler) GetMonthlyTransactionMethodSuccess(ctx context.Context, req *pb.MonthMethodTransactionRequest) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	data, err := h.repo.GetMonthlyTransactionMethodSuccess(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionMonthPaymentMethod{
		Status:  "success",
		Message: "success get monthly transaction method success",
		Data:    data,
	}, nil
}

func (h *TransactionStatsHandler) GetYearlyTransactionMethodSuccess(ctx context.Context, req *pb.YearMethodTransactionRequest) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	data, err := h.repo.GetYearlyTransactionMethodSuccess(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionYearPaymentmethod{
		Status:  "success",
		Message: "success get yearly transaction method success",
		Data:    data,
	}, nil
}

func (h *TransactionStatsHandler) GetMonthlyTransactionMethodFailed(ctx context.Context, req *pb.MonthMethodTransactionRequest) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	data, err := h.repo.GetMonthlyTransactionMethodFailed(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionMonthPaymentMethod{
		Status:  "success",
		Message: "success get monthly transaction method failed",
		Data:    data,
	}, nil
}

func (h *TransactionStatsHandler) GetYearlyTransactionMethodFailed(ctx context.Context, req *pb.YearMethodTransactionRequest) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	data, err := h.repo.GetYearlyTransactionMethodFailed(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionYearPaymentmethod{
		Status:  "success",
		Message: "success get yearly transaction method failed",
		Data:    data,
	}, nil
}

type TransactionStatsByMerchantHandler struct {
	pb.UnimplementedTransactionStatsByMerchantServiceServer
	repo repository.Repository
	log  logger.LoggerInterface
}

func NewTransactionStatsByMerchantHandler(repo repository.Repository, log logger.LoggerInterface) *TransactionStatsByMerchantHandler {
	return &TransactionStatsByMerchantHandler{
		repo: repo,
		log:  log,
	}
}

func (h *TransactionStatsByMerchantHandler) GetMonthlyAmountSuccessByMerchant(ctx context.Context, req *pb.MonthAmountTransactionMerchantRequest) (*pb.ApiResponseTransactionMonthAmountSuccess, error) {
	data, err := h.repo.GetMonthlyAmountSuccessByMerchant(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionMonthAmountSuccess{
		Status:  "success",
		Message: "success get monthly amount success by merchant",
		Data:    data,
	}, nil
}

func (h *TransactionStatsByMerchantHandler) GetYearlyAmountSuccessByMerchant(ctx context.Context, req *pb.YearAmountTransactionMerchantRequest) (*pb.ApiResponseTransactionYearAmountSuccess, error) {
	data, err := h.repo.GetYearlyAmountSuccessByMerchant(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionYearAmountSuccess{
		Status:  "success",
		Message: "success get yearly amount success by merchant",
		Data:    data,
	}, nil
}

func (h *TransactionStatsByMerchantHandler) GetMonthlyAmountFailedByMerchant(ctx context.Context, req *pb.MonthAmountTransactionMerchantRequest) (*pb.ApiResponseTransactionMonthAmountFailed, error) {
	data, err := h.repo.GetMonthlyAmountFailedByMerchant(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionMonthAmountFailed{
		Status:  "success",
		Message: "success get monthly amount failed by merchant",
		Data:    data,
	}, nil
}

func (h *TransactionStatsByMerchantHandler) GetYearlyAmountFailedByMerchant(ctx context.Context, req *pb.YearAmountTransactionMerchantRequest) (*pb.ApiResponseTransactionYearAmountFailed, error) {
	data, err := h.repo.GetYearlyAmountFailedByMerchant(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionYearAmountFailed{
		Status:  "success",
		Message: "success get yearly amount failed by merchant",
		Data:    data,
	}, nil
}

func (h *TransactionStatsByMerchantHandler) GetMonthlyTransactionMethodByMerchantSuccess(ctx context.Context, req *pb.MonthMethodTransactionMerchantRequest) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	data, err := h.repo.GetMonthlyTransactionMethodByMerchantSuccess(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionMonthPaymentMethod{
		Status:  "success",
		Message: "success get monthly transaction method success by merchant",
		Data:    data,
	}, nil
}

func (h *TransactionStatsByMerchantHandler) GetYearlyTransactionMethodByMerchantSuccess(ctx context.Context, req *pb.YearMethodTransactionMerchantRequest) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	data, err := h.repo.GetYearlyTransactionMethodByMerchantSuccess(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionYearPaymentmethod{
		Status:  "success",
		Message: "success get yearly transaction method success by merchant",
		Data:    data,
	}, nil
}

func (h *TransactionStatsByMerchantHandler) GetMonthlyTransactionMethodByMerchantFailed(ctx context.Context, req *pb.MonthMethodTransactionMerchantRequest) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	data, err := h.repo.GetMonthlyTransactionMethodByMerchantFailed(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionMonthPaymentMethod{
		Status:  "success",
		Message: "success get monthly transaction method failed by merchant",
		Data:    data,
	}, nil
}

func (h *TransactionStatsByMerchantHandler) GetYearlyTransactionMethodByMerchantFailed(ctx context.Context, req *pb.YearMethodTransactionMerchantRequest) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	data, err := h.repo.GetYearlyTransactionMethodByMerchantFailed(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseTransactionYearPaymentmethod{
		Status:  "success",
		Message: "success get yearly transaction method failed by merchant",
		Data:    data,
	}, nil
}

package handler

import (
	"context"
	"github.com/MamangRust/monolith-ecommerce-pkg/logger"
	"github.com/MamangRust/monolith-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-ecommerce-grpc/service/stats-reader/repository"
)

type OrderStatsHandler struct {
	pb.UnimplementedOrderStatsServiceServer
	repo repository.Repository
	log  logger.LoggerInterface
}

func NewOrderStatsHandler(repo repository.Repository, log logger.LoggerInterface) *OrderStatsHandler {
	return &OrderStatsHandler{
		repo: repo,
		log:  log,
	}
}

func (h *OrderStatsHandler) FindMonthlyTotalRevenue(ctx context.Context, req *pb.FindYearMonthTotalRevenue) (*pb.ApiResponseOrderMonthlyTotalRevenue, error) {
	data, err := h.repo.FindMonthlyTotalRevenue(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseOrderMonthlyTotalRevenue{
		Status:  "success",
		Message: "success get monthly total revenue",
		Data:    data,
	}, nil
}

func (h *OrderStatsHandler) FindYearlyTotalRevenue(ctx context.Context, req *pb.FindYearTotalRevenue) (*pb.ApiResponseOrderYearlyTotalRevenue, error) {
	data, err := h.repo.FindYearlyTotalRevenue(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseOrderYearlyTotalRevenue{
		Status:  "success",
		Message: "success get yearly total revenue",
		Data:    data,
	}, nil
}

func (h *OrderStatsHandler) FindMonthlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearMonthTotalRevenueByMerchant) (*pb.ApiResponseOrderMonthlyTotalRevenue, error) {
	data, err := h.repo.FindMonthlyTotalRevenueByMerchant(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseOrderMonthlyTotalRevenue{
		Status:  "success",
		Message: "success get monthly total revenue by merchant",
		Data:    data,
	}, nil
}

func (h *OrderStatsHandler) FindYearlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearTotalRevenueByMerchant) (*pb.ApiResponseOrderYearlyTotalRevenue, error) {
	data, err := h.repo.FindYearlyTotalRevenueByMerchant(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseOrderYearlyTotalRevenue{
		Status:  "success",
		Message: "success get yearly total revenue by merchant",
		Data:    data,
	}, nil
}

func (h *OrderStatsHandler) FindMonthlyRevenue(ctx context.Context, req *pb.FindYearOrder) (*pb.ApiResponseOrderMonthly, error) {
	data, err := h.repo.FindMonthlyRevenue(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseOrderMonthly{
		Status:  "success",
		Message: "success get monthly order revenue",
		Data:    data,
	}, nil
}

func (h *OrderStatsHandler) FindYearlyRevenue(ctx context.Context, req *pb.FindYearOrder) (*pb.ApiResponseOrderYearly, error) {
	data, err := h.repo.FindYearlyRevenue(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseOrderYearly{
		Status:  "success",
		Message: "success get yearly order revenue",
		Data:    data,
	}, nil
}

func (h *OrderStatsHandler) FindMonthlyRevenueByMerchant(ctx context.Context, req *pb.FindYearOrderByMerchant) (*pb.ApiResponseOrderMonthly, error) {
	data, err := h.repo.FindMonthlyRevenueByMerchant(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseOrderMonthly{
		Status:  "success",
		Message: "success get monthly order revenue by merchant",
		Data:    data,
	}, nil
}

func (h *OrderStatsHandler) FindYearlyRevenueByMerchant(ctx context.Context, req *pb.FindYearOrderByMerchant) (*pb.ApiResponseOrderYearly, error) {
	data, err := h.repo.FindYearlyRevenueByMerchant(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseOrderYearly{
		Status:  "success",
		Message: "success get yearly order revenue by merchant",
		Data:    data,
	}, nil
}

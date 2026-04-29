package handler

import (
	"context"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/microservice-ecommerce-grpc/service/stats-reader/repository"
)

type CategoryStatsHandler struct {
	pb.UnimplementedCategoryStatsServiceServer
	repo repository.Repository
	log  logger.LoggerInterface
}

func NewCategoryStatsHandler(repo repository.Repository, log logger.LoggerInterface) *CategoryStatsHandler {
	return &CategoryStatsHandler{
		repo: repo,
		log:  log,
	}
}

func (h *CategoryStatsHandler) FindMonthlyTotalPrices(ctx context.Context, req *pb.FindYearMonthTotalPrices) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
	data, err := h.repo.FindMonthlyTotalPrices(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseCategoryMonthlyTotalPrice{
		Status:  "success",
		Message: "success get monthly total prices",
		Data:    data,
	}, nil
}

func (h *CategoryStatsHandler) FindYearlyTotalPrices(ctx context.Context, req *pb.FindYearTotalPrices) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	data, err := h.repo.FindYearlyTotalPrices(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseCategoryYearlyTotalPrice{
		Status:  "success",
		Message: "success get yearly total prices",
		Data:    data,
	}, nil
}

func (h *CategoryStatsHandler) FindMonthPrice(ctx context.Context, req *pb.FindYearCategory) (*pb.ApiResponseCategoryMonthPrice, error) {
	data, err := h.repo.FindMonthPrice(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseCategoryMonthPrice{
		Status:  "success",
		Message: "success get monthly category prices",
		Data:    data,
	}, nil
}

func (h *CategoryStatsHandler) FindYearPrice(ctx context.Context, req *pb.FindYearCategory) (*pb.ApiResponseCategoryYearPrice, error) {
	data, err := h.repo.FindYearPrice(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseCategoryYearPrice{
		Status:  "success",
		Message: "success get yearly category prices",
		Data:    data,
	}, nil
}

type CategoryStatsByMerchantHandler struct {
	pb.UnimplementedCategoryStatsByMerchantServiceServer
	repo repository.Repository
	log  logger.LoggerInterface
}

func NewCategoryStatsByMerchantHandler(repo repository.Repository, log logger.LoggerInterface) *CategoryStatsByMerchantHandler {
	return &CategoryStatsByMerchantHandler{
		repo: repo,
		log:  log,
	}
}

func (h *CategoryStatsByMerchantHandler) FindMonthlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearMonthTotalPriceByMerchant) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
	data, err := h.repo.FindMonthlyTotalPricesByMerchant(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseCategoryMonthlyTotalPrice{
		Status:  "success",
		Message: "success get monthly total prices by merchant",
		Data:    data,
	}, nil
}

func (h *CategoryStatsByMerchantHandler) FindYearlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearTotalPriceByMerchant) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	data, err := h.repo.FindYearlyTotalPricesByMerchant(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseCategoryYearlyTotalPrice{
		Status:  "success",
		Message: "success get yearly total prices by merchant",
		Data:    data,
	}, nil
}

func (h *CategoryStatsByMerchantHandler) FindMonthPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) (*pb.ApiResponseCategoryMonthPrice, error) {
	data, err := h.repo.FindMonthPriceByMerchant(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseCategoryMonthPrice{
		Status:  "success",
		Message: "success get monthly category prices by merchant",
		Data:    data,
	}, nil
}

func (h *CategoryStatsByMerchantHandler) FindYearPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) (*pb.ApiResponseCategoryYearPrice, error) {
	data, err := h.repo.FindYearPriceByMerchant(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseCategoryYearPrice{
		Status:  "success",
		Message: "success get yearly category prices by merchant",
		Data:    data,
	}, nil
}

type CategoryStatsByIdHandler struct {
	pb.UnimplementedCategoryStatsByIdServiceServer
	repo repository.Repository
	log  logger.LoggerInterface
}

func NewCategoryStatsByIdHandler(repo repository.Repository, log logger.LoggerInterface) *CategoryStatsByIdHandler {
	return &CategoryStatsByIdHandler{
		repo: repo,
		log:  log,
	}
}

func (h *CategoryStatsByIdHandler) FindMonthlyTotalPricesById(ctx context.Context, req *pb.FindYearMonthTotalPriceById) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
	data, err := h.repo.FindMonthlyTotalPricesById(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseCategoryMonthlyTotalPrice{
		Status:  "success",
		Message: "success get monthly total prices by id",
		Data:    data,
	}, nil
}

func (h *CategoryStatsByIdHandler) FindYearlyTotalPricesById(ctx context.Context, req *pb.FindYearTotalPriceById) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	data, err := h.repo.FindYearlyTotalPricesById(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseCategoryYearlyTotalPrice{
		Status:  "success",
		Message: "success get yearly total prices by id",
		Data:    data,
	}, nil
}

func (h *CategoryStatsByIdHandler) FindMonthPriceById(ctx context.Context, req *pb.FindYearCategoryById) (*pb.ApiResponseCategoryMonthPrice, error) {
	data, err := h.repo.FindMonthPriceById(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseCategoryMonthPrice{
		Status:  "success",
		Message: "success get monthly category prices by id",
		Data:    data,
	}, nil
}

func (h *CategoryStatsByIdHandler) FindYearPriceById(ctx context.Context, req *pb.FindYearCategoryById) (*pb.ApiResponseCategoryYearPrice, error) {
	data, err := h.repo.FindYearPriceById(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ApiResponseCategoryYearPrice{
		Status:  "success",
		Message: "success get yearly category prices by id",
		Data:    data,
	}, nil
}

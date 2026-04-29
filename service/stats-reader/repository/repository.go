package repository

import (
	"context"
	"github.com/MamangRust/monolith-ecommerce-shared/pb"
)

type Repository interface {
	// Category Stats
	FindMonthlyTotalPrices(ctx context.Context, req *pb.FindYearMonthTotalPrices) ([]*pb.CategoriesMonthlyTotalPriceResponse, error)
	FindYearlyTotalPrices(ctx context.Context, req *pb.FindYearTotalPrices) ([]*pb.CategoriesYearlyTotalPriceResponse, error)
	FindMonthPrice(ctx context.Context, req *pb.FindYearCategory) ([]*pb.CategoryMonthPriceResponse, error)
	FindYearPrice(ctx context.Context, req *pb.FindYearCategory) ([]*pb.CategoryYearPriceResponse, error)

	// Category Stats By ID
	FindMonthlyTotalPricesById(ctx context.Context, req *pb.FindYearMonthTotalPriceById) ([]*pb.CategoriesMonthlyTotalPriceResponse, error)
	FindYearlyTotalPricesById(ctx context.Context, req *pb.FindYearTotalPriceById) ([]*pb.CategoriesYearlyTotalPriceResponse, error)
	FindMonthPriceById(ctx context.Context, req *pb.FindYearCategoryById) ([]*pb.CategoryMonthPriceResponse, error)
	FindYearPriceById(ctx context.Context, req *pb.FindYearCategoryById) ([]*pb.CategoryYearPriceResponse, error)

	// Category Stats By Merchant
	FindMonthlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearMonthTotalPriceByMerchant) ([]*pb.CategoriesMonthlyTotalPriceResponse, error)
	FindYearlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearTotalPriceByMerchant) ([]*pb.CategoriesYearlyTotalPriceResponse, error)
	FindMonthPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) ([]*pb.CategoryMonthPriceResponse, error)
	FindYearPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) ([]*pb.CategoryYearPriceResponse, error)

	// Order Stats
	FindMonthlyTotalRevenue(ctx context.Context, req *pb.FindYearMonthTotalRevenue) ([]*pb.OrderMonthlyTotalRevenueResponse, error)
	FindYearlyTotalRevenue(ctx context.Context, req *pb.FindYearTotalRevenue) ([]*pb.OrderYearlyTotalRevenueResponse, error)
	FindMonthlyRevenue(ctx context.Context, req *pb.FindYearOrder) ([]*pb.OrderMonthlyResponse, error)
	FindYearlyRevenue(ctx context.Context, req *pb.FindYearOrder) ([]*pb.OrderYearlyResponse, error)

	// Order Stats By Merchant
	FindMonthlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearMonthTotalRevenueByMerchant) ([]*pb.OrderMonthlyTotalRevenueResponse, error)
	FindYearlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearTotalRevenueByMerchant) ([]*pb.OrderYearlyTotalRevenueResponse, error)
	FindMonthlyRevenueByMerchant(ctx context.Context, req *pb.FindYearOrderByMerchant) ([]*pb.OrderMonthlyResponse, error)
	FindYearlyRevenueByMerchant(ctx context.Context, req *pb.FindYearOrderByMerchant) ([]*pb.OrderYearlyResponse, error)

	// Transaction Stats
	GetMonthlyAmountSuccess(ctx context.Context, req *pb.MonthAmountTransactionRequest) ([]*pb.TransactionMonthlyAmountSuccess, error)
	GetYearlyAmountSuccess(ctx context.Context, req *pb.YearAmountTransactionRequest) ([]*pb.TransactionYearlyAmountSuccess, error)
	GetMonthlyAmountFailed(ctx context.Context, req *pb.MonthAmountTransactionRequest) ([]*pb.TransactionMonthlyAmountFailed, error)
	GetYearlyAmountFailed(ctx context.Context, req *pb.YearAmountTransactionRequest) ([]*pb.TransactionYearlyAmountFailed, error)
	GetMonthlyTransactionMethodSuccess(ctx context.Context, req *pb.MonthMethodTransactionRequest) ([]*pb.TransactionMonthlyMethod, error)
	GetYearlyTransactionMethodSuccess(ctx context.Context, req *pb.YearMethodTransactionRequest) ([]*pb.TransactionYearlyMethod, error)
	GetMonthlyTransactionMethodFailed(ctx context.Context, req *pb.MonthMethodTransactionRequest) ([]*pb.TransactionMonthlyMethod, error)
	GetYearlyTransactionMethodFailed(ctx context.Context, req *pb.YearMethodTransactionRequest) ([]*pb.TransactionYearlyMethod, error)

	// Transaction Stats By Merchant
	GetMonthlyAmountSuccessByMerchant(ctx context.Context, req *pb.MonthAmountTransactionMerchantRequest) ([]*pb.TransactionMonthlyAmountSuccess, error)
	GetYearlyAmountSuccessByMerchant(ctx context.Context, req *pb.YearAmountTransactionMerchantRequest) ([]*pb.TransactionYearlyAmountSuccess, error)
	GetMonthlyAmountFailedByMerchant(ctx context.Context, req *pb.MonthAmountTransactionMerchantRequest) ([]*pb.TransactionMonthlyAmountFailed, error)
	GetYearlyAmountFailedByMerchant(ctx context.Context, req *pb.YearAmountTransactionMerchantRequest) ([]*pb.TransactionYearlyAmountFailed, error)
	GetMonthlyTransactionMethodByMerchantSuccess(ctx context.Context, req *pb.MonthMethodTransactionMerchantRequest) ([]*pb.TransactionMonthlyMethod, error)
	GetYearlyTransactionMethodByMerchantSuccess(ctx context.Context, req *pb.YearMethodTransactionMerchantRequest) ([]*pb.TransactionYearlyMethod, error)
	GetMonthlyTransactionMethodByMerchantFailed(ctx context.Context, req *pb.MonthMethodTransactionMerchantRequest) ([]*pb.TransactionMonthlyMethod, error)
	GetYearlyTransactionMethodByMerchantFailed(ctx context.Context, req *pb.YearMethodTransactionMerchantRequest) ([]*pb.TransactionYearlyMethod, error)
}

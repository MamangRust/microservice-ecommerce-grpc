package analytics_test

import (
	"context"
	"testing"

	"github.com/MamangRust/monolith-ecommerce-grpc/service/stats-reader/handler"
	"github.com/MamangRust/monolith-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-ecommerce-pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

// MockReaderRepository is a mock implementation of the repository interface
type MockReaderRepository struct {
	mock.Mock
}

func (m *MockReaderRepository) FindMonthlyTotalPrices(ctx context.Context, req *pb.FindYearMonthTotalPrices) ([]*pb.CategoriesMonthlyTotalPriceResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.CategoriesMonthlyTotalPriceResponse), args.Error(1)
}

func (m *MockReaderRepository) FindYearlyTotalPrices(ctx context.Context, req *pb.FindYearTotalPrices) ([]*pb.CategoriesYearlyTotalPriceResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.CategoriesYearlyTotalPriceResponse), args.Error(1)
}

func (m *MockReaderRepository) FindMonthPrice(ctx context.Context, req *pb.FindYearCategory) ([]*pb.CategoryMonthPriceResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.CategoryMonthPriceResponse), args.Error(1)
}

func (m *MockReaderRepository) FindYearPrice(ctx context.Context, req *pb.FindYearCategory) ([]*pb.CategoryYearPriceResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.CategoryYearPriceResponse), args.Error(1)
}

// ... other methods omitted for brevity, adding only what's used in these tests

func (m *MockReaderRepository) FindMonthlyTotalPricesById(ctx context.Context, req *pb.FindYearMonthTotalPriceById) ([]*pb.CategoriesMonthlyTotalPriceResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.CategoriesMonthlyTotalPriceResponse), args.Error(1)
}
func (m *MockReaderRepository) FindYearlyTotalPricesById(ctx context.Context, req *pb.FindYearTotalPriceById) ([]*pb.CategoriesYearlyTotalPriceResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.CategoriesYearlyTotalPriceResponse), args.Error(1)
}
func (m *MockReaderRepository) FindMonthPriceById(ctx context.Context, req *pb.FindYearCategoryById) ([]*pb.CategoryMonthPriceResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.CategoryMonthPriceResponse), args.Error(1)
}
func (m *MockReaderRepository) FindYearPriceById(ctx context.Context, req *pb.FindYearCategoryById) ([]*pb.CategoryYearPriceResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.CategoryYearPriceResponse), args.Error(1)
}
func (m *MockReaderRepository) FindMonthlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearMonthTotalPriceByMerchant) ([]*pb.CategoriesMonthlyTotalPriceResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.CategoriesMonthlyTotalPriceResponse), args.Error(1)
}
func (m *MockReaderRepository) FindYearlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearTotalPriceByMerchant) ([]*pb.CategoriesYearlyTotalPriceResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.CategoriesYearlyTotalPriceResponse), args.Error(1)
}
func (m *MockReaderRepository) FindMonthPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) ([]*pb.CategoryMonthPriceResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.CategoryMonthPriceResponse), args.Error(1)
}
func (m *MockReaderRepository) FindYearPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) ([]*pb.CategoryYearPriceResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.CategoryYearPriceResponse), args.Error(1)
}
func (m *MockReaderRepository) FindMonthlyTotalRevenue(ctx context.Context, req *pb.FindYearMonthTotalRevenue) ([]*pb.OrderMonthlyTotalRevenueResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.OrderMonthlyTotalRevenueResponse), args.Error(1)
}
func (m *MockReaderRepository) FindYearlyTotalRevenue(ctx context.Context, req *pb.FindYearTotalRevenue) ([]*pb.OrderYearlyTotalRevenueResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.OrderYearlyTotalRevenueResponse), args.Error(1)
}
func (m *MockReaderRepository) FindMonthlyRevenue(ctx context.Context, req *pb.FindYearOrder) ([]*pb.OrderMonthlyResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.OrderMonthlyResponse), args.Error(1)
}
func (m *MockReaderRepository) FindYearlyRevenue(ctx context.Context, req *pb.FindYearOrder) ([]*pb.OrderYearlyResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.OrderYearlyResponse), args.Error(1)
}
func (m *MockReaderRepository) FindMonthlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearMonthTotalRevenueByMerchant) ([]*pb.OrderMonthlyTotalRevenueResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.OrderMonthlyTotalRevenueResponse), args.Error(1)
}
func (m *MockReaderRepository) FindYearlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearTotalRevenueByMerchant) ([]*pb.OrderYearlyTotalRevenueResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.OrderYearlyTotalRevenueResponse), args.Error(1)
}
func (m *MockReaderRepository) FindMonthlyRevenueByMerchant(ctx context.Context, req *pb.FindYearOrderByMerchant) ([]*pb.OrderMonthlyResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.OrderMonthlyResponse), args.Error(1)
}
func (m *MockReaderRepository) FindYearlyRevenueByMerchant(ctx context.Context, req *pb.FindYearOrderByMerchant) ([]*pb.OrderYearlyResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.OrderYearlyResponse), args.Error(1)
}
func (m *MockReaderRepository) GetMonthlyAmountSuccess(ctx context.Context, req *pb.MonthAmountTransactionRequest) ([]*pb.TransactionMonthlyAmountSuccess, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionMonthlyAmountSuccess), args.Error(1)
}
func (m *MockReaderRepository) GetYearlyAmountSuccess(ctx context.Context, req *pb.YearAmountTransactionRequest) ([]*pb.TransactionYearlyAmountSuccess, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionYearlyAmountSuccess), args.Error(1)
}
func (m *MockReaderRepository) GetMonthlyAmountFailed(ctx context.Context, req *pb.MonthAmountTransactionRequest) ([]*pb.TransactionMonthlyAmountFailed, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionMonthlyAmountFailed), args.Error(1)
}
func (m *MockReaderRepository) GetYearlyAmountFailed(ctx context.Context, req *pb.YearAmountTransactionRequest) ([]*pb.TransactionYearlyAmountFailed, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionYearlyAmountFailed), args.Error(1)
}
func (m *MockReaderRepository) GetMonthlyTransactionMethodSuccess(ctx context.Context, req *pb.MonthMethodTransactionRequest) ([]*pb.TransactionMonthlyMethod, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionMonthlyMethod), args.Error(1)
}
func (m *MockReaderRepository) GetYearlyTransactionMethodSuccess(ctx context.Context, req *pb.YearMethodTransactionRequest) ([]*pb.TransactionYearlyMethod, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionYearlyMethod), args.Error(1)
}
func (m *MockReaderRepository) GetMonthlyTransactionMethodFailed(ctx context.Context, req *pb.MonthMethodTransactionRequest) ([]*pb.TransactionMonthlyMethod, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionMonthlyMethod), args.Error(1)
}
func (m *MockReaderRepository) GetYearlyTransactionMethodFailed(ctx context.Context, req *pb.YearMethodTransactionRequest) ([]*pb.TransactionYearlyMethod, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionYearlyMethod), args.Error(1)
}
func (m *MockReaderRepository) GetMonthlyAmountSuccessByMerchant(ctx context.Context, req *pb.MonthAmountTransactionMerchantRequest) ([]*pb.TransactionMonthlyAmountSuccess, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionMonthlyAmountSuccess), args.Error(1)
}
func (m *MockReaderRepository) GetYearlyAmountSuccessByMerchant(ctx context.Context, req *pb.YearAmountTransactionMerchantRequest) ([]*pb.TransactionYearlyAmountSuccess, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionYearlyAmountSuccess), args.Error(1)
}
func (m *MockReaderRepository) GetMonthlyAmountFailedByMerchant(ctx context.Context, req *pb.MonthAmountTransactionMerchantRequest) ([]*pb.TransactionMonthlyAmountFailed, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionMonthlyAmountFailed), args.Error(1)
}
func (m *MockReaderRepository) GetYearlyAmountFailedByMerchant(ctx context.Context, req *pb.YearAmountTransactionMerchantRequest) ([]*pb.TransactionYearlyAmountFailed, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionYearlyAmountFailed), args.Error(1)
}
func (m *MockReaderRepository) GetMonthlyTransactionMethodByMerchantSuccess(ctx context.Context, req *pb.MonthMethodTransactionMerchantRequest) ([]*pb.TransactionMonthlyMethod, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionMonthlyMethod), args.Error(1)
}
func (m *MockReaderRepository) GetYearlyTransactionMethodByMerchantSuccess(ctx context.Context, req *pb.YearMethodTransactionMerchantRequest) ([]*pb.TransactionYearlyMethod, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionYearlyMethod), args.Error(1)
}
func (m *MockReaderRepository) GetMonthlyTransactionMethodByMerchantFailed(ctx context.Context, req *pb.MonthMethodTransactionMerchantRequest) ([]*pb.TransactionMonthlyMethod, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionMonthlyMethod), args.Error(1)
}
func (m *MockReaderRepository) GetYearlyTransactionMethodByMerchantFailed(ctx context.Context, req *pb.YearMethodTransactionMerchantRequest) ([]*pb.TransactionYearlyMethod, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*pb.TransactionYearlyMethod), args.Error(1)
}

func TestCategoryStatsHandler_FindMonthlyTotalPrices(t *testing.T) {
	mockRepo := new(MockReaderRepository)
	z, _ := zap.NewDevelopment()
	l := &logger.Logger{Log: z}
	h := handler.NewCategoryStatsHandler(mockRepo, l)

	req := &pb.FindYearMonthTotalPrices{Year: 2024, Month: 1}
	expectedData := []*pb.CategoriesMonthlyTotalPriceResponse{{Year: "2024", Month: "January", TotalRevenue: 1000}}
	
	mockRepo.On("FindMonthlyTotalPrices", mock.Anything, req).Return(expectedData, nil)

	res, err := h.FindMonthlyTotalPrices(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "success", res.Status)
	assert.Equal(t, expectedData, res.Data)
	mockRepo.AssertExpectations(t)
}

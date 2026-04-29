package analytics_test

import (
	"context"
	"testing"
	"time"

	"github.com/MamangRust/microservice-ecommerce-grpc/service/stats-writer/usecase"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockWriterRepository is a mock implementation of the repository interface
type MockWriterRepository struct {
	mock.Mock
}

func (m *MockWriterRepository) InsertCategoryStat(ctx context.Context, event events.CategoryStatEvent) error {
	args := m.Called(ctx, event)
	return args.Error(0)
}

func (m *MockWriterRepository) InsertOrderStat(ctx context.Context, event events.OrderStatEvent) error {
	args := m.Called(ctx, event)
	return args.Error(0)
}

func (m *MockWriterRepository) InsertTransactionStat(ctx context.Context, event events.TransactionStatEvent) error {
	args := m.Called(ctx, event)
	return args.Error(0)
}

func TestStatsUseCase_SaveCategoryStat(t *testing.T) {
	mockRepo := new(MockWriterRepository)
	uc := usecase.NewStatsUseCase(mockRepo)

	event := events.CategoryStatEvent{
		CategoryID:   1,
		CategoryName: "Test Category",
		TotalViews:   10,
		CreatedAt:    time.Now(),
	}

	mockRepo.On("InsertCategoryStat", mock.Anything, event).Return(nil)

	err := uc.SaveCategoryStat(context.Background(), event)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStatsUseCase_SaveOrderStat(t *testing.T) {
	mockRepo := new(MockWriterRepository)
	uc := usecase.NewStatsUseCase(mockRepo)

	event := events.OrderStatEvent{
		OrderID:            1,
		MerchantID:         200,
		TotalRevenue:       50000,
		TotalItemsSold:     2,
		CreatedAt:          time.Now(),
	}

	mockRepo.On("InsertOrderStat", mock.Anything, event).Return(nil)

	err := uc.SaveOrderStat(context.Background(), event)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStatsUseCase_SaveTransactionStat(t *testing.T) {
	mockRepo := new(MockWriterRepository)
	uc := usecase.NewStatsUseCase(mockRepo)

	event := events.TransactionStatEvent{
		TransactionID: 1,
		OrderID:       10,
		UserID:        100,
		MerchantID:    200,
		Amount:        50000,
		PaymentMethod: "credit_card",
		Status:        "success",
		CreatedAt:     time.Now(),
	}

	mockRepo.On("InsertTransactionStat", mock.Anything, event).Return(nil)

	err := uc.SaveTransactionStat(context.Background(), event)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

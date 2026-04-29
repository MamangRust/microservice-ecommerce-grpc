package handler

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-grpc/service/stats-writer/usecase"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/events"
	"go.uber.org/zap"
)

type StatsHandler struct {
	useCase usecase.UseCase
	log     logger.LoggerInterface
}

func NewStatsHandler(useCase usecase.UseCase, log logger.LoggerInterface) *StatsHandler {
	return &StatsHandler{
		useCase: useCase,
		log:     log,
	}
}

func (h *StatsHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *StatsHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h *StatsHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		switch msg.Topic {
		case "category_stats_topic":
			var event events.CategoryStatEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				h.log.Error("Failed to unmarshal category stat event", zap.Error(err))
				continue
			}
			if err := h.useCase.SaveCategoryStat(context.Background(), event); err != nil {
				h.log.Error("Failed to save category stat event", zap.Error(err))
				continue
			}
		case "order_stats_topic":
			var event events.OrderStatEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				h.log.Error("Failed to unmarshal order stat event", zap.Error(err))
				continue
			}
			if err := h.useCase.SaveOrderStat(context.Background(), event); err != nil {
				h.log.Error("Failed to save order stat event", zap.Error(err))
				continue
			}
		case "transaction_stats_topic":
			var event events.TransactionStatEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				h.log.Error("Failed to unmarshal transaction stat event", zap.Error(err))
				continue
			}
			if err := h.useCase.SaveTransactionStat(context.Background(), event); err != nil {
				h.log.Error("Failed to save transaction stat event", zap.Error(err))
				continue
			}
		}

		session.MarkMessage(msg, "")
	}
	return nil
}

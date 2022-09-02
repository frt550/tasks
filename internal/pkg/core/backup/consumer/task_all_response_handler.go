package consumer

import (
	"encoding/json"
	"tasks/internal/pkg/core/logger"
	"tasks/pkg/contract/kafka"
	"time"

	"github.com/Shopify/sarama"
	"github.com/patrickmn/go-cache"
)

type handler struct {
	cache *cache.Cache
}

func New(cache *cache.Cache) sarama.ConsumerGroupHandler {
	return &handler{
		cache: cache,
	}
}

func (h *handler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *handler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case <-session.Context().Done():
			logger.Logger.Info("task_all_response_handler: done")
			return nil
		case msg, ok := <-claim.Messages():
			if !ok {
				logger.Logger.Info("task_all_response_handler: data channel closed")
				return nil
			}

			message := &kafka.TaskAllResponseMessage{}
			err := json.Unmarshal(msg.Value, message)
			if err != nil {
				logger.Logger.Sugar().Error(err)
			}
			logger.Logger.Sugar().Infof("task_all_response_handler: received new message, request id %s", message.RequestId)
			h.cache.Set(message.RequestId, message.Data, 24*time.Hour)

			session.MarkMessage(msg, "")
		}
	}
}

package consumer

import (
	"context"
	"encoding/json"
	"tasks/internal/pkg/core/logger"
	producerPkg "tasks/internal/pkg/core/producer"
	taskPkg "tasks/internal/pkg/core/task"
	"tasks/pkg/contract/kafka"

	"github.com/Shopify/sarama"
)

type handler struct {
	producer    producerPkg.Interface
	taskService taskPkg.Interface
}

func New(producer producerPkg.Interface, taskService taskPkg.Interface) sarama.ConsumerGroupHandler {
	return &handler{
		producer:    producer,
		taskService: taskService,
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
			logger.Logger.Info("task_all_request_handler: done")
			return nil
		case msg, ok := <-claim.Messages():
			if !ok {
				logger.Logger.Info("task_all_request_handler: data channel closed")
				return nil
			}

			message := &kafka.TaskAllRequestMessage{}
			err := json.Unmarshal(msg.Value, message)
			if err != nil {
				logger.Logger.Sugar().Error(err)
			}

			logger.Logger.Sugar().Infof("task_all_request_handler: received new message, request id %s", message.RequestId)

			allTasks, err := h.taskService.All(context.Background(), 0, 0)
			j, err := json.Marshal(allTasks)
			if err != nil {
				logger.Logger.Sugar().Error(err)
				return err
			}

			err = h.producer.Publish(kafka.TopicTaskAllResponse, &kafka.TaskAllResponseMessage{
				RequestId: message.RequestId,
				Data:      string(j),
			})
			if err != nil {
				logger.Logger.Sugar().Error(err)
				return err
			}

			session.MarkMessage(msg, "")
		}
	}
}

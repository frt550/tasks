package producer

import (
	"encoding/json"
	"tasks/internal/pkg/core/logger"

	"github.com/Shopify/sarama"
)

type Interface interface {
	Publish(topic string, message interface{}) error
}

func New(brokers []string, config *sarama.Config) Interface {
	syncProducer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		logger.Logger.Sugar().Fatal(err)
	}
	return &producer{syncProducer: syncProducer}
}

type producer struct {
	syncProducer sarama.SyncProducer
}

func (p *producer) Publish(topic string, message interface{}) error {
	j, err := json.Marshal(message)
	if err != nil {
		logger.Logger.Sugar().Error(err)
		return err
	}
	_, _, err = p.syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder("sync"),
		Value: sarama.ByteEncoder(j),
	})
	if err != nil {
		logger.Logger.Sugar().Error(err)
		return err
	} else {
		logger.Logger.Sugar().Infof("Message '%s' is sent to '%s' topic", j, topic)
		return nil
	}
}

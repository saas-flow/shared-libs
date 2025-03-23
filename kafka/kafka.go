package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/fx"
)

var ProducerModule = fx.Module("kafka.producer",
	fx.Provide(NewProducer),
)

func NewProducer(config *kafka.ConfigMap) (*kafka.Producer, error) {
	return kafka.NewProducer(config)
}

var ConsumerModule = fx.Module("kafka.consumer",
	fx.Provide(NewConsumer),
)

func NewConsumer(config *kafka.ConfigMap) (*kafka.Consumer, error) {
	return kafka.NewConsumer(config)
}

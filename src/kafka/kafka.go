package kafka

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewProducer(kafkaServer string) (*kafka.Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
	})
}

func NewMessage(dto interface{}, topic string) *kafka.Message {

	value, err := json.Marshal(dto)
	if err != nil {
		panic(err)
	}

	return &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}
}

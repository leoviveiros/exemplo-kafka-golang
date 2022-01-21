package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	producer := NewKafkaProducer()

	Publish("Hello World", "teste", producer, nil)
	producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "exemplo-kafka-golang_kafka_1:9092",
	}

	producer, err := kafka.NewProducer(configMap)

	if err != nil {
		panic(err)
	}

	return producer
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte) error {
	message := &kafka.Message{
		Value: []byte(msg),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key: key,
	}

	err := producer.Produce(message, nil)

	if err != nil {
		return err
	}

	return nil
}
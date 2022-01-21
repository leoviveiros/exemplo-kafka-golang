package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {	
	producer := NewKafkaProducer()
	deliveryChan := make(chan kafka.Event)

	Publish("Hello World", "teste", producer, nil, deliveryChan)
	
	event := <-deliveryChan
	msg := event.(*kafka.Message)
	
	if msg.TopicPartition.Error != nil {
		fmt.Println("Erro ao enviar mensagem: ", msg.TopicPartition.Error)
	} else {
		fmt.Println("Mensagem enviada - partition: ", msg.TopicPartition.Partition, " offset: ", msg.TopicPartition.Offset)
	}

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

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value: []byte(msg),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key: key,
	}

	err := producer.Produce(message, deliveryChan)

	if err != nil {
		return err
	}

	return nil
}
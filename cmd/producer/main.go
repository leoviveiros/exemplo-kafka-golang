package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {	
	producer := NewKafkaProducer()
	deliveryChan := make(chan kafka.Event)

	Publish("Hello World", "teste", producer, nil, deliveryChan)
	
	go DeliveyReport(deliveryChan)

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

func DeliveyReport(deliveryChan chan kafka.Event) {
	for event := range deliveryChan {
		switch ev := event.(type) {
		case *kafka.Message: 
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar mensagem: ", ev.TopicPartition.Error)
			} else {
				fmt.Println("Mensagem enviada - partition: ", ev.TopicPartition.Partition, " offset: ", ev.TopicPartition.Offset)
			}
		}
	}
}
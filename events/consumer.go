package events

import (
	"context"
	"log"
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type ConsumerConfig struct {
	Brokers string
	Topic   string
	GroupID string
}

func getKafkaReader(kafkaUrl string, topic string, groupId string) *kafka.Reader {
	brokers := strings.Split(kafkaUrl, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupId,
		Topic:          topic,
		MinBytes:       1,        // 1B
		MaxBytes:       57671680, // 5MB
		MaxWait:        100 * time.Millisecond,
		StartOffset:    kafka.LastOffset,
		SessionTimeout: 15 * time.Second,
	})
}

func RunConsumer(messageHandler func(message *kafka.Message) error, consumerConfig ConsumerConfig) {
	// get kafka reader
	reader := getKafkaReader(consumerConfig.Brokers, consumerConfig.Topic, consumerConfig.GroupID)
	defer reader.Close()

	// looping to wait and read message
	log.Println("Run Consumer for topic", consumerConfig.Topic)
	for {
		// recieve message from broker
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
		}

		handlerErr := messageHandler(&message)
		if handlerErr != nil {
			log.Println(handlerErr)
		}
	}
}

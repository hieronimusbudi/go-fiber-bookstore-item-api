package events

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type Message struct {
	Subject string      `json:"subject"`
	Context interface{} `json:"context"`
}

type MessageContextExample struct {
	ID       string `json:"id"`
	Field    string `json:"field"`
	Previous string `json:"previous"`
	New      string `json:"new"`
}

// {
// 	"action": "user_updated",
// 	"context": {
// 	  "id": 1343,
// 	  "field": "first_name",
// 	  "previous": "Robert",
// 	  "new": "Bob"
// 	}
// }

// func onCompletion(message []kafka.Message, err error) {
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	log.Println("message delivery success!")
// }

type ProducerConfig struct {
	Addr       string
	Topic      string
	Completion func(message []kafka.Message, err error)
}

func getKafkaWriter(addr, topic string, completionFunction func(message []kafka.Message, err error)) *kafka.Writer {
	kafkaWriter := &kafka.Writer{
		Addr:         kafka.TCP(addr),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    1,
		BatchTimeout: 10 * time.Millisecond,
	}

	if completionFunction != nil {
		kafkaWriter.Completion = completionFunction
	}

	return kafkaWriter
}

func producerHandler(kafkaWriter *kafka.Writer, message *Message, key string) {
	jsonOut, jsonErr := json.Marshal(message)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	msg := kafka.Message{
		Value: []byte(fmt.Sprint(string(jsonOut))),
	}
	if key != "" {
		msg.Key = []byte(key)
	}

	err := kafkaWriter.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Println(err)
	}
}

func RunProducer(message *Message, key string, producerConfig ProducerConfig) {
	// get kafka writer
	kafkaWriter := getKafkaWriter(producerConfig.Addr, producerConfig.Topic, producerConfig.Completion)
	defer kafkaWriter.Close()

	// send message
	producerHandler(kafkaWriter, message, key)

	log.Println("Run Producer for topic", producerConfig.Topic)
}

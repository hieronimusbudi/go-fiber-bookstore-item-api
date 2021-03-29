package events

import (
	"log"

	eventsutils "github.com/hieronimusbudi/go-bookstore-utils/events"
	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/events"
	"github.com/segmentio/kafka-go"
)

func onCompletionItemDeletedEvent(message []kafka.Message, err error) {
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("event delivery success!")
}

func ProduceItemDeletedEvent(itemId string) {
	itemID, idErr := primitive.ObjectIDFromHex(itemId)
	if idErr != nil {
		log.Println(resterrors.NewInternalServerError("create object id error", idErr))
		return
	}

	message := eventsutils.Message{
		Subject: eventsutils.ItemUpdated,
		Context: eventsutils.ItemCreatedContext{
			ID: itemID,
		},
	}

	eventsutils.RunProducer(&message, "", eventsutils.ProducerConfig{
		Addr:       events.KafkaURLLocal,
		Topic:      events.TopicItemDeletedLocal,
		Completion: onCompletionItemDeletedEvent,
	})
}

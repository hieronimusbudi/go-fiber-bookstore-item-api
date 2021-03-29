package events

import (
	"log"

	eventsutils "github.com/hieronimusbudi/go-bookstore-utils/events"

	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/domain/items"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/events"
	"github.com/segmentio/kafka-go"
)

func onCompletionItemUpdatedEvent(message []kafka.Message, err error) {
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("event delivery success!")
}

func ProduceItemUpdatedEvent(item *items.Item) {
	message := eventsutils.Message{
		Subject: eventsutils.ItemUpdated,
		Context: eventsutils.ItemCreatedContext{
			ID:                item.ID,
			Seller:            item.Seller,
			Title:             item.Title,
			Price:             item.Price,
			AvailableQuantity: item.AvailableQuantity,
			SoldQuantity:      item.SoldQuantity,
			Status:            item.Status,
		},
	}

	eventsutils.RunProducer(&message, "", eventsutils.ProducerConfig{
		Addr:       events.KafkaURLLocal,
		Topic:      events.TopicItemUpdatedLocal,
		Completion: onCompletionItemUpdatedEvent,
	})
}

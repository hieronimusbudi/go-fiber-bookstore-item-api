package events

import envvar "github.com/hieronimusbudi/go-fiber-bookstore-item-api/env"

var (
	KafkaURL      = envvar.KafkaURL
	KafkaURLLocal = "localhost:9092"

	TopicItemCreated      = envvar.KafkaTopic
	TopicItemCreatedLocal = "test-item-created"

	TopicItemUpdated      = envvar.KafkaTopic
	TopicItemUpdatedLocal = "test-item-updated"

	TopicItemDeleted      = envvar.KafkaTopic
	TopicItemDeletedLocal = "test-item-deleted"
)

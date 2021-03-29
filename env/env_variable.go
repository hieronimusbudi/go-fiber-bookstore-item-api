package envvar

import "os"

var (
	MongoURI         = os.Getenv("MONGO_URI")
	ElasticSearchURI = os.Getenv("ELASTICS_URI")
	KafkaURL         = os.Getenv("KAFKA_URL")
	KafkaTopic       = os.Getenv("KAFKA_TOPIC")
	KafkaGroupID     = os.Getenv("KAFKA_GROUP_ID")
	JwtSecret        = os.Getenv("JWT_SECRET")
	JwtCookieName    = os.Getenv("JWT_COOKIE_NAME")
)

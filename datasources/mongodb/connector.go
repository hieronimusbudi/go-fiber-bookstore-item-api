package mongodb

import (
	"context"
	"log"

	envvar "github.com/hieronimusbudi/go-fiber-bookstore-item-api/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ItemsCollection *mongo.Collection
	mongoURI        = envvar.MongoURI
	mongoURILocal   = "mongodb://localhost:27017/items"
)

func Init() {
	clientOptions := options.Client().ApplyURI(mongoURILocal)

	client, clientErr := mongo.Connect(context.TODO(), clientOptions)
	if clientErr != nil {
		panic(clientErr)
	}

	pingErr := client.Ping(context.TODO(), nil)
	if pingErr != nil {
		panic(pingErr)
	}

	ItemsCollection = client.Database("items").Collection("items")
	log.Println("Connected to MongoDB!")
}

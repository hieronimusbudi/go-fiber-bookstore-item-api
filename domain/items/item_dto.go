package items

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	ESID              string             `json:"es_id" bson:"es_id" `
	Seller            int64              `json:"seller" bson:"seller"`
	Title             string             `json:"title" bson:"title"`
	Description       Description        `json:"description" bson:"description"`
	Pictures          []Picture          `json:"pictures" bson:"pictures"`
	Video             string             `json:"video" bson:"video"`
	Price             float32            `json:"price" bson:"price"`
	AvailableQuantity int                `json:"available_quantity" bson:"available_quantity"`
	SoldQuantity      int                `json:"sold_quantity" bson:"sold_quantity"`
	Status            string             `json:"status" bson:"status"`
}

type Description struct {
	PlainText string `json:"plain_text" bson:"plain_text"`
	Html      string `json:"html" bson:"html"`
}

type Picture struct {
	ID  int64  `json:"id" bson:"id"`
	URL string `json:"url" bson:"url"`
}

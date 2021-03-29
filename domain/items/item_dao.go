package items

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/datasources/elasticsearch"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/datasources/mongodb"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/domain/queries"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() resterrors.RestErr {
	// saving in elastic
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)

	if err != nil {
		return resterrors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.ESID = result.Id

	// saving in mongodb
	i.ID = primitive.NewObjectID()
	insertResult, insertErr := mongodb.ItemsCollection.InsertOne(context.TODO(), i)
	if insertErr != nil {
		return resterrors.NewInternalServerError("error when trying to save item", insertErr)
	}

	itemId, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return resterrors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.ID = itemId

	return nil
}

func (i *Item) Get() resterrors.RestErr {
	itemId := i.ID
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.ID.Hex())

	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return resterrors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.ID))
		}
		return resterrors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.ID), errors.New("database error"))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}

	if err := json.Unmarshal(bytes, &i); err != nil {
		return resterrors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}
	i.ID = itemId
	return nil
}

func (i *Item) GetSingle() resterrors.RestErr {
	// itemID, idErr := primitive.ObjectIDFromHex(i.ID)
	// if idErr != nil {
	// 	return resterrors.NewInternalServerError("create object id error", idErr)
	// }

	filter := bson.M{"_id": i.ID}
	findErr := mongodb.ItemsCollection.FindOne(context.TODO(), filter).Decode(&i)
	if findErr != nil {
		if strings.Contains(findErr.Error(), "404") {
			return resterrors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.ID))
		}
		return resterrors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.ID), errors.New(findErr.Error()))
	}

	return nil
}

func (i *Item) Update() resterrors.RestErr {
	filter := bson.M{"_id": i.ID}
	iByte, err := bson.Marshal(*i)
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to marshal item", err)
	}

	var update bson.M
	err = bson.Unmarshal(iByte, &update)
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to nnmarshal item", err)
	}

	_, updateErr := mongodb.ItemsCollection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: update}})
	if updateErr != nil {
		return resterrors.NewInternalServerError("error when trying to update item", updateErr)
	}

	return nil
}

func (i *Item) Delete() resterrors.RestErr {
	filter := bson.M{"_id": i.ID}

	_, deleteErr := mongodb.ItemsCollection.DeleteOne(context.TODO(), filter)
	if deleteErr != nil {
		return resterrors.NewInternalServerError("error when trying to delete item", deleteErr)
	}

	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, resterrors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, resterrors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, resterrors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		item.ESID = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, resterrors.NewNotFoundError("no items found matching given criteria")
	}
	return items, nil
}

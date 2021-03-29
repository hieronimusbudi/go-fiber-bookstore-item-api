package services

import (
	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/domain/items"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/domain/queries"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type itemsService struct{}

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, resterrors.RestErr)
	Get(string) (*items.Item, resterrors.RestErr)
	Update(string, items.Item) (*items.Item, resterrors.RestErr)
	Delete(string) resterrors.RestErr
	Search(queries.EsQuery) ([]items.Item, resterrors.RestErr)
}

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

func (s *itemsService) Create(item items.Item) (*items.Item, resterrors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, resterrors.RestErr) {
	itemID, idErr := primitive.ObjectIDFromHex(id)
	if idErr != nil {
		return nil, resterrors.NewInternalServerError("create object id error", idErr)
	}
	item := items.Item{ID: itemID}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Update(id string, item items.Item) (*items.Item, resterrors.RestErr) {
	itemID, idErr := primitive.ObjectIDFromHex(id)
	if idErr != nil {
		return nil, resterrors.NewInternalServerError("create object id error", idErr)
	}

	item.ID = itemID
	if err := item.Update(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Delete(id string) resterrors.RestErr {
	itemID, idErr := primitive.ObjectIDFromHex(id)
	if idErr != nil {
		return resterrors.NewInternalServerError("create object id error", idErr)
	}

	item := items.Item{ID: itemID}
	if err := item.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, resterrors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}

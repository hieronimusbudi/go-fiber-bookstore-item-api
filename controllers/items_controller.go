package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/domain/items"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/domain/queries"
	events "github.com/hieronimusbudi/go-fiber-bookstore-item-api/events/producer"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/services"
)

type itemsController struct{}
type itemsControllerInterface interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Search(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

func (ctr *itemsController) Create(c *fiber.Ctx) error {
	itemRequest := new(items.Item)
	if err := c.BodyParser(itemRequest); err != nil {
		respErr := resterrors.NewBadRequestError("invalid item json body")
		return c.Status(400).JSON(respErr)
	}

	itemRequest.Seller = 1
	result, createErr := services.ItemsService.Create(*itemRequest)
	if createErr != nil {
		return c.Status(400).JSON(createErr)
	}

	go events.ProduceItemCreatedEvent(result)

	return c.Status(http.StatusCreated).JSON(result)
}

func (ctr *itemsController) Get(c *fiber.Ctx) error {
	itemId := c.Params("id")

	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(item)
}

func (ctr *itemsController) Update(c *fiber.Ctx) error {
	itemId := c.Params("id")

	itemRequest := new(items.Item)
	if err := c.BodyParser(itemRequest); err != nil {
		respErr := resterrors.NewBadRequestError("invalid item json body")
		return c.Status(400).JSON(respErr)
	}

	item, err := services.ItemsService.Update(itemId, *itemRequest)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	go events.ProduceItemUpdatedEvent(item)

	return c.Status(http.StatusOK).JSON(item)
}

func (ctr *itemsController) Delete(c *fiber.Ctx) error {
	itemId := c.Params("id")

	err := services.ItemsService.Delete(itemId)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	go events.ProduceItemDeletedEvent(itemId)

	return c.Status(http.StatusOK).JSON("delete success")
}

func (ctr *itemsController) Search(c *fiber.Ctx) error {
	query := new(queries.EsQuery)
	if err := c.BodyParser(query); err != nil {
		err := resterrors.NewBadRequestError("invalid json body")
		return c.Status(400).JSON(err)
	}

	items, searchErr := services.ItemsService.Search(*query)
	if searchErr != nil {
		return c.Status(400).JSON(searchErr)
	}

	return c.Status(http.StatusOK).JSON(items)
}

package elasticsearch

import (
	"context"
	"fmt"
	"log"
	"time"

	envvar "github.com/hieronimusbudi/go-fiber-bookstore-item-api/env"

	"github.com/hieronimusbudi/go-bookstore-utils/logger"
	"github.com/olivere/elastic"
)

var (
	host      = envvar.ElasticSearchURI
	hostLocal = "http://localhost:9200"
)

type esClient struct {
	client *elastic.Client
}

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
}

var (
	Client esClientInterface = &esClient{}
)

func Init() {
	myLog := logger.GetLogger()

	client, err := elastic.NewClient(
		elastic.SetURL(hostLocal),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(myLog),
		elastic.SetInfoLog(myLog),
	)
	if err != nil {
		log.Println(err)
	}

	Client.setClient(client)
	log.Println("Connected to ElasticSearch")
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	log.Println("hello", index, docType)
	result, err := c.client.Index().Index(index).Type(docType).BodyJson(doc).Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().Index(index).Type(docType).Id(id).Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get id %s", id), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := c.client.Search(index).
		Query(query).
		RestTotalHitsAsInt(true).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search documents in index %s", index), err)
		return nil, err
	}
	return result, nil
}

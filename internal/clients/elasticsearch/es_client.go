package elasticsearch

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
	CreateIndex(string, string) error
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)

	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()

	result, err := c.client.Index().
		Index(index).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		message := fmt.Sprintf("error when trying to index document in index %s.\nError: %+v", index, err)
		log.Println(message, err)
		return nil, err
	}

	return result, nil
}

func (c *esClient) CreateIndex(name string, body string) error {
	ctx := context.Background()

	createIndex, err := c.client.CreateIndex(name).BodyString(body).Do(ctx)
	if err != nil {
		return err
	}

	if !createIndex.Acknowledged {
		return errors.New("index not acknowledge")
	}

	return nil
}

package main

import (
	"github.com/wgarcia4190/bookstore_items_api/internal/clients/elasticsearch"
)

func main() {
	elasticsearch.Init()

	mapping := `{
		"settings": {
			"index": {
				"number_of_shards": 4,
				"number_of_replicas": 2
			}
		}
	}`

	if err := elasticsearch.Client.CreateIndex("items", mapping); err != nil {
		panic(err)
	}
}

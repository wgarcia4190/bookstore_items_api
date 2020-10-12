package items

import (
	"errors"

	"github.com/wgarcia4190/bookstore_items_api/internal/clients/elasticsearch"
	"github.com/wgarcia4190/bookstore_utils_go/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}

	i.ID = result.Id

	return nil
}

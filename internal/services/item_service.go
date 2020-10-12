package services

import (
	"github.com/wgarcia4190/bookstore_items_api/internal/domain/items"
	"github.com/wgarcia4190/bookstore_utils_go/rest_errors"
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemService struct{}

var (
	ItemsService itemsServiceInterface = &itemService{}
)

func (s *itemService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Get(id string) (*items.Item, *rest_errors.RestErr) {
	return nil, nil
}

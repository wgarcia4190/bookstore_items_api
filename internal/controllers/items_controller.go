package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wgarcia4190/bookstore_utils_go/rest_errors"

	"github.com/wgarcia4190/bookstore_items_api/internal/utils"

	"github.com/wgarcia4190/bookstore_items_api/internal/domain/items"
	"github.com/wgarcia4190/bookstore_items_api/internal/services"
	"github.com/wgarcia4190/bookstore_oauth_go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid request body")
		utils.RespondError(w, restErr)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid item json body")
		utils.RespondError(w, restErr)
		return
	}

	itemRequest.Seller = oauth.GetCallerId(r)

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		utils.RespondError(w, createErr)
	}

	utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}

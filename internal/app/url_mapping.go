package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wgarcia4190/bookstore_items_api/internal/controllers"
)

func mapUrl(router *mux.Router) {
	router.HandleFunc("/health", controllers.HealthController.Health).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}

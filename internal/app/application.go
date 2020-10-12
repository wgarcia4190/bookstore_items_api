package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/wgarcia4190/bookstore_items_api/internal/clients/elasticsearch"
)

func StartApplication() {
	elasticsearch.Init()

	router := mux.NewRouter()
	mapUrl(router)

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  6 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}

package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ddo/go-mux-mvc/controllers/product"
	"github.com/ddo/go-mux-mvc/controllers/web"
)

// New .
func New() http.Handler {
	r := mux.NewRouter()

	// website
	r.HandleFunc("/", web.Root).Methods("GET")

	// product
	r.HandleFunc("/product", product.New).Methods("POST")
	r.HandleFunc("/product/{id}", product.Get).Methods("GET")

	return r
}

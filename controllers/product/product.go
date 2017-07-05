package product

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ddo/go-mux-mvc/models/product"
)

// New .
func New(w http.ResponseWriter, r *http.Request) {
	_product, err := product.New(r.FormValue("name"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	fmt.Fprintf(w, "%s", _product.ID.Hex())
	return
}

// Get .
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_product, err := product.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if _product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, _product.String())
	return
}

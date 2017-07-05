package web

import (
	"fmt"
	"net/http"
)

// Root .
func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}

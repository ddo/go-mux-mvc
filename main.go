package main

import (
	"net/http"
	"os"

	"github.com/ddo/go-mux-mvc/controllers"
	"github.com/ddo/go-mux-mvc/models/logger"

	// init db
	_ "github.com/ddo/go-mux-mvc/db/mongodb"
	_ "github.com/ddo/go-mux-mvc/setting"
)

const (
	defaultPort = "8008"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// route
	handler := controllers.New()

	logger.Log("port:", port)
	err := http.ListenAndServe("0.0.0.0:"+port, handler)
	if err != nil {
		logger.Log("ERR ListenAndServe:", err)
	}
}

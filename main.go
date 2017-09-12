package main

import (
	"net/http"
	"os"
	"time"

	"github.com/ddo/go-mux-mvc/controllers"
	"github.com/ddo/go-mux-mvc/models/logger"

	// init db
	_ "github.com/ddo/go-mux-mvc/db/mongodb"
	_ "github.com/ddo/go-mux-mvc/setting"
)

const (
	defaultPort = "8008"

	idleTimeout       = 30 * time.Second
	writeTimeout      = 180 * time.Second
	readHeaderTimeout = 10 * time.Second
	readTimeout       = 10 * time.Second
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	logger.Log("port:", port)

	// route
	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: handler,

		IdleTimeout:       idleTimeout,
		WriteTimeout:      writeTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Log("ERR ListenAndServe:", err)
	}
}

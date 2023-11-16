package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
)

func main() {
	dbConnection, err := managers.InitializeDB()
	if err != nil {
		// TODO: Log proper information
	}
	defer dbConnection.Close()

	router := createRouter(dbConnection)

	const port = "8080"

	server := &http.Server{
		Addr:              "localhost:" + port,
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// TODO: Log proper information
		}
	}()

	<-quit

	if err := server.Shutdown(context.TODO()); err != nil {
		// TODO: Log proper information
	}

	os.Exit(0)
}

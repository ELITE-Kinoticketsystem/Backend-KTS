package main
import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
)


func main() {
	log.Println("Initializing database connection...")
	dbConnection, err := managers.InitializeDB()
	if err != nil {
		panic(err)
	}
	log.Println("Database initialized successfully")
	defer dbConnection.Close()

	router := createRouter(dbConnection)

	const port = "8080"

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		log.Printf("Server listening on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error starting server: %v\n", err)
		}
	}()

	<-quit

	if err := server.Shutdown(context.TODO()); err != nil {
		log.Fatalf("error shutting down server: %v\n", err)
	}

	os.Exit(0)
	
}

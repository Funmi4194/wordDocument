package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Funmi4194/myMod/database"
	"github.com/Funmi4194/myMod/routes"
	"github.com/gorilla/mux"
)

func main() {
	//setup the mux router
	h := mux.NewRouter()

	//connect to database
	database.ConnectDB()

	//call the endpoints into the handler
	routes.UserRoute(h)

	s := &http.Server{
		Addr:    ":80",
		Handler: h,
	}

	//create a new goroutines to start the server
	go func() {
		log.Printf("Starting server on http://localhost:%d", 80)
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("error listening on port: %s\n", err)
		}
	}()

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt)
	<-osSignal
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shut down...")
	}
	log.Println("Server exited!")
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Funmi4194/myMod/config"
	"github.com/Funmi4194/myMod/database"
	"github.com/Funmi4194/myMod/middleware"
	"github.com/Funmi4194/myMod/routes"
	"github.com/Funmi4194/myMod/types"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func createServer() (s *http.Server) {
	//setup the mux router
	h := mux.NewRouter()

	//wrappping router into custom recover middleware
	r := middleware.Recover(h)

	//injest combined logger
	logger := handlers.CombinedLoggingHandler(os.Stdout, r)

	//call the endpoints into the handler
	routes.UserRoute(h)

	//load .env files
	env := config.MustGet("ENV_PATH", ".env")
	log.Printf("Loading %s file\n", env)
	if err := godotenv.Load(env); err != nil {
		if err := godotenv.Load(); err != nil {
			log.Printf("Error verifying environment varaiables: %s\n", err)
		}
	}

	//verify env variables
	if err := config.VerifyEnvironment(types.Env{}); err != nil {
		log.Fatalf("Error verifying environment variables: %s\n", err)
	}

	//append env variables to config.Env
	config.AppendEnvironment(config.Env)

	//connect to database
	client, err := database.NewMongoDBClient(config.Env.DocumentDBUri)
	if err != nil {
		log.Fatalf("Error connecting to DocumentDB: %s\n", err)
	}

	//select database
	database.WordCountDB = client.Database(config.Env.DocumentDB)

	port := fmt.Sprintf(":%s", config.Env.Port)

	s = &http.Server{
		Addr:    port,
		Handler: logger,
	}

	//create a new goroutines to start the server
	go func() {
		log.Printf("Starting server on http://localhost%s", port)
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("error listening on port: %s\n", err)
		}
	}()
	return s
}

func main() {
	s := createServer()

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

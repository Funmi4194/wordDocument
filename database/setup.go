package database

import (
	"context"
	"log"
	"time"

	//"github.com/Funmi4194/myMod/config"
	"github.com/Funmi4194/myMod/config"
	"github.com/Funmi4194/myMod/environ"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() *mongo.Client {
	environ.Environ()
	clientOptions := options.Client().ApplyURI(config.Env.DocumentDBUri)
	//fmt.Print(config.Env.DocumentDBUri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Connected to MongoDB!")
	return client
}

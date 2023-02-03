package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb+srv://funmi-tech:funmi4194@cluster0.v5gihly.mongodb.net/?retryWrites=true&w=majority")

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
	fmt.Println("Connected to MongoDB!")
	return client

}

var DB *mongo.Client = ConnectDB()

func GetConnection(collectionInstance string) *mongo.Collection {
	collection := DB.Database("WordCount").Collection(collectionInstance)
	return collection
}

package database

import "go.mongodb.org/mongo-driver/mongo"

var WordCountDB *mongo.Database

func GetConnection(collectionInstance string) *mongo.Collection {
	WordCountDB = ConnectDB().Database("WordCount")
	collection := WordCountDB.Collection(collectionInstance)
	return collection
}

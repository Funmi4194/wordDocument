package database

import (
	"github.com/Funmi4194/myMod/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var WordCountDB *mongo.Database

func GetConnection(collectionInstance string) *mongo.Collection {
	WordCountDB = ConnectDB().Database(config.Env.DocumentDB)
	collection := WordCountDB.Collection(collectionInstance)
	return collection

}

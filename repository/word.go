package repo

import (
	"context"
	"log"

	"github.com/Funmi4194/myMod/config"
	"github.com/Funmi4194/myMod/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create a new document
func (w *WordCount) Create() error {
	//insert a new document
	//w takes all the values in the struct WordCount
	result, err := database.WordCountDB.Collection(config.WordCollection).InsertOne(context.Background(), w)
	if err != nil {
		return err
	}
	//displays the inserted struct if sucessful
	w.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

//Find documents by documentName
func (w *WordCount) FindDocument() error {
	//create a find filter
	filter := bson.D{primitive.E{Key: "document-name", Value: w.DocumentName}}
	if err := database.WordCountDB.Collection(config.WordCollection).FindOne(context.Background(), filter).Decode(&w); err != nil {

		return err

	}
	return nil
}

func (w *WordCounts) Documents() error {
	result, err := database.WordCountDB.Collection(config.WordCollection).Find(context.Background(), bson.D{})
	if err != nil {
		return err
	} else {
		if err = result.All(context.Background(), w); err != nil {
			log.Fatal(err)
		}

	}
	return nil
}

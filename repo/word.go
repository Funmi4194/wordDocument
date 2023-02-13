package repo

import (
	"context"

	"github.com/Funmi4194/myMod/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = database.GetConnection("WordDocuments")

//Create a new document
func (w *WordCount) Create() error {
	//insert a new document
	//w takes all the values in the struct WordCount
	result, err := UserCollection.InsertOne(context.Background(), w)
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
	filter := bson.D{primitive.E{Key: "DocumentName", Value: w.DocumentName}}
	if err := UserCollection.FindOne(context.Background(), filter).Decode(&w); err != nil {
		//if err == mongo.ErrNoDocuments {
		return err
		//}
	}
	return nil
}

//Find all documents
func (w *WordCount) FindDocuments() error {
	result, err := UserCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	} else {
		if err := result.All(context.Background(), w); err != nil {
			return err
		}
	}

	return nil
}

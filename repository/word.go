package repo

import (
	"context"

	"github.com/Funmi4194/myMod/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//var UserCollection mongo.Collection = database.WordCountDB.Collection("WordDocuments")

//var UserCollection *mongo.Collection = database.GetConnection("WordDocuments")

//Create a new document
func (w *WordCount) Create() error {
	//insert a new document
	//w takes all the values in the struct WordCount
	result, err := database.WordCountDB.Collection("WordDocuments").InsertOne(context.Background(), w)
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
	if err := database.WordCountDB.Collection("WordDocuments").FindOne(context.Background(), filter).Decode(&w); err != nil {

		return err

	}
	return nil
}

// //Find all documents
// func (w *WordCount) Documents() (WordCounts, error) {
// 	result, err := UserCollection.Find(context.Background(), bson.D{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	var documents WordCounts
// 	defer result.Close(context.Background())
// 	for result.Next(context.Background()) {
// 		var document WordCount
// 		err = result.Decode(&document)
// 		if err != nil {
// 			return nil, err
// 		}
// 		documents = append(documents, document)
// 	}
// 	if err := result.Err(); err != nil {
// 		return nil, err
// 	}

// 	return documents, nil
// }

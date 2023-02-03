package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Funmi4194/myMod/configs"
	"github.com/Funmi4194/myMod/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var usercollection *mongo.Collection = configs.GetConnection("WordDocuments")
var document models.WordCount

func CreateDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewDecoder(r.Body).Decode(&document); err != nil {
			log.Fatal(err)
		}

		content := strings.Join(strings.Fields(document.Content), " ")
		totalCount := len(strings.Split(content, " "))

		totalCharacters := len(strings.Split(document.Content, ""))

		//(total character without space)
		content2 := strings.Replace(document.Content, " ", "", -1)
		charWithoutSpace := len(strings.Split(content2, ""))

		//numbers of sentences
		sentences := len(strings.Split(document.Content, "."))

		// numbers of paragraphs
		paragraphs := len(strings.Split(document.Content, "\n\n"))

		//numbers of line

		lines := strings.SplitAfterN(document.Content, "\n", -1)
		line := len(lines)

		documents := models.WordCount{
			DocumentName:     strings.ToLower(document.DocumentName),
			Content:          document.Content,
			Words:            totalCount,
			Characters:       totalCharacters,
			CharWithoutSpace: charWithoutSpace,
			Sentence:         sentences,
			Paragraphs:       paragraphs,
			Lines:            line,
		}

		result, err := usercollection.InsertOne(context.Background(), documents)
		document.ID = result.InsertedID.(primitive.ObjectID)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(documents)
	}
}

func GetDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var document models.WordCount
		params := mux.Vars(r)
		doc := strings.ToLower(params["doc"])

		filter := bson.M{"DocumentName": doc}
		err := usercollection.FindOne(context.Background(), filter).Decode(&document)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(document)
	}
}

func GetDocuments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, err := usercollection.Find(context.Background(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		var documents []models.WordCount
		defer result.Close(context.Background())

		for result.Next(context.Background()) {
			var document models.WordCount
			err = result.Decode(&document)
			if err != nil {
				log.Fatal(err)
			}
			documents = append(documents, document)
		}
		json.NewEncoder(w).Encode(documents)

	}

}

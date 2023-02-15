package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Funmi4194/myMod/logic"
	"github.com/Funmi4194/myMod/repo"
	response "github.com/Funmi4194/myMod/responses"
	"github.com/gorilla/mux"
)

//CreateDocument creates a new document history and returns an handlerFunc
func CreateDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		var document repo.WordCount
		//Decode json objects into the body
		if err := json.NewDecoder(r.Body).Decode(&document); err != nil {
			response.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
			return
		}

		content := document.Content
		wordName := document.DocumentName
		// create word takes wordname and content from the document struct instance
		documents, err := logic.CreateWord(wordName, content)
		if err != nil {

			response.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
			return

		}

		response.SendJSONResponse(w, true, http.StatusOK, "Created sucessfully", response.M{"word": documents})
	}
}

func GetDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// set variable params
		params := mux.Vars(r)
		doc := strings.ToLower(params["doc"])

		documents, err := logic.GetDocument(doc)
		if err != nil {

			response.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
			// return

		}

		response.SendJSONResponse(w, true, http.StatusOK, "Document details retrieved", response.M{"word": documents})

	}
}

//var usercollection *mongo.Collection = database.GetConnection("WordDocuments")

// func GetDocuments() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")

// 		result, err := usercollection.Find(context.Background(), bson.M{})
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		var documents []repo.WordCount
// 		defer result.Close(context.Background())

// 		for result.Next(context.Background()) {
// 			var document repo.WordCount
// 			err = result.Decode(&document)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			documents = append(documents, document)
// 		}
// 		json.NewEncoder(w).Encode(documents)

// 	}

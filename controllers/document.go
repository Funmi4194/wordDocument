package controllers

import (
	"context"
	"net/http"

	"github.com/Funmi4194/myMod/logic"
	repo "github.com/Funmi4194/myMod/repository"
	response "github.com/Funmi4194/myMod/responses"
	"github.com/Funmi4194/myMod/types"
)

func CreateDocument(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	body := r.Context().Value(types.BodyCtxKey{}).(*repo.WordCount)

	documents, err := logic.CreateWord(*body)
	if err != nil {

		response.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
		return

	}
	ctx := context.WithValue(r.Context(), types.AuthCtxKey{}, &documents.DocumentName)

	// load request with context
	r = r.WithContext(ctx)

	response.SendJSONResponse(w, true, http.StatusOK, "Created sucessfully", response.M{"word": documents}, r)
}

func GetDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// set variable params
	params := r.Context().Value(types.ParamsCtxKey{}).(map[string]string)

	//set doc to GetDocument to search for document
	documents, err := logic.GetDocument(params["doc"])
	if err != nil {

		response.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
		return

	}

	response.SendJSONResponse(w, true, http.StatusOK, "Document details retrieved", response.M{"word": documents})

}

func GetDocuments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		document, err := logic.GetDocuments()

		if err != nil {

			response.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
			return

		}

		response.SendJSONResponse(w, true, http.StatusOK, "Document details retrieved", response.M{"word": document})
	}
}

// for result.Next(context.Background()) {
// 	var document WordCount
// 	err = result.Decode(&document)
// 	if err != nil {
// 		return nil, err
// 	}
// 	documents = append(documents, document)
// }
// if err := result.Err(); err != nil {
// 	return nil, err
// }
// }

// defer display.Close(context.Background())
// for display.Next(context.Background()) {
// 	//used bson.D because order matters to me
// 	var user bson.D
// 	err = display.Decode(&user)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(user)
// }

// }

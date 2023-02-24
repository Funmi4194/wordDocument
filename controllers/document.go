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

	//add the body to the logic function
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

func GetDocuments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	document, err := logic.GetDocuments()

	if err != nil {

		response.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
		return

	}

	response.SendJSONResponse(w, true, http.StatusOK, "Document details retrieved", response.M{"word": document})
}

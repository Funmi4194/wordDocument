package routes

import (
	"github.com/Funmi4194/myMod/controllers"
	"github.com/Funmi4194/myMod/middleware"
	repo "github.com/Funmi4194/myMod/repository"
	"github.com/gorilla/mux"
)

func UserRoute(route *mux.Router) {

	//handler for creating document
	route.Handle("/createDoc", middleware.BodyWithFunc(&repo.WordCount{})(controllers.CreateDocument)).Methods("POST")

	//handler to get document
	route.Handle("/getDoc/{doc}", middleware.ParamsWithFunc()(controllers.GetDocument)).Methods("GET")

	route.HandleFunc("/getAllDoc", (controllers.GetDocuments)).Methods("GET")

}

package routes

import (
	"github.com/Funmi4194/myMod/controllers"
	"github.com/gorilla/mux"
)

func UserRoute(route *mux.Router) {
	//handlerfunc for creating document
	route.HandleFunc("/createDoc", controllers.CreateDocument()).Methods("POST")

	//handlerfunc to get document
	route.HandleFunc("/getDoc/{doc}", controllers.GetDocument()).Methods("GET")
	//route.HandleFunc("/getAllDoc", controllers.GetDocuments()).Methods("GET")

}

package routes

import (
	"github.com/Funmi4194/myMod/controllers"
	"github.com/gorilla/mux"
)

func UserRoute(route *mux.Router) {

	route.HandleFunc("/createDoc", controllers.CreateDocument()).Methods("POST")
	route.HandleFunc("/getDoc/{doc}", controllers.GetDocument()).Methods("GET")
	//route.HandleFunc("/getAllDoc", controllers.GetDocuments()).Methods("GET")

}

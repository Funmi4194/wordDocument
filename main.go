package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Funmi4194/myMod/configs"
	"github.com/Funmi4194/myMod/routes"
	"github.com/gorilla/mux"
)

func main() {
	h := mux.NewRouter()

	configs.ConnectDB()
	routes.UserRoute(h)
	s := &http.Server{
		Addr:    ":80",
		Handler: h,
	}
	fmt.Printf("Starting server on http://localhost:%d\n", 80)
	log.Fatal(s.ListenAndServe())
}

package main

import (
	"log"
	"net/http"

	n "velociraptorgo/datastore"
	"velociraptorgo/handler"

	"github.com/gorilla/mux"
)

func main() {
	n.LoadDataStore()
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api").Subrouter()
	// Route configuration
	sub.Methods("GET").Path("/gettotalvelociraptor/{timestamp}").HandlerFunc(handler.GetTotalVelociraptor)
	sub.Methods("POST").Path("/updatevelociraptor").HandlerFunc(handler.UpdateVelociraptor)
	log.Fatal(http.ListenAndServe(":3000", router))
}

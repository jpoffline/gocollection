package collectionapi

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/collections", CollectionNames).Methods("GET")
	myRouter.HandleFunc("/fields/{collectionname}", ReturnFieldNames).Methods("GET")
	myRouter.HandleFunc("/field/{collectionname}", UpdateFieldName).Methods("PATCH")
	myRouter.HandleFunc("/collection/{collectionname}", ReturnCollectionAsTable).Methods("GET")
	myRouter.HandleFunc("/collection/{collectionname}", AddRecordToCollection).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", handlers.CORS()(myRouter)))
}

package collectionapi

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/collections", CollectionNames)
	myRouter.HandleFunc("/collection/{collectionname}", ReturnCollectionAsTable)
	log.Fatal(http.ListenAndServe(":10000", handlers.CORS()(myRouter)))
}

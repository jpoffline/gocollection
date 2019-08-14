package collectionapi

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/collections", CollectionNames).Methods("GET")
	myRouter.HandleFunc("/collection/{collectionname}", ReturnCollectionAsTable).Methods("GET")
	myRouter.HandleFunc("/collection/{collectionname}", AddRecordToCollection).Methods("POST")
	myRouter.HandleFunc("/fields/{collectionname}", ReturnFieldNames).Methods("GET")
	myRouter.HandleFunc("/field/{collectionname}", UpdateFieldName).Methods("PATCH")
	myRouter.HandleFunc("/meta/{collectionname}", ReturnCollectionMeta).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(originsOk, headersOk, methodsOk)(myRouter)))
}

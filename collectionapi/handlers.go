package collectionapi

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "X-Custom-Header"})
	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	corsObj := handlers.AllowedOrigins([]string{"*"})

	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "UPDATE"})

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/collections", CollectionNames).Methods("GET")
	myRouter.HandleFunc("/collection/{collectionname}", ReturnCollectionAsTable).Methods("GET")
	myRouter.HandleFunc("/collection/{collectionname}", AddRecordToCollection).Methods("POST")
	myRouter.HandleFunc("/fields/{collectionname}", ReturnFieldNames).Methods("GET")
	myRouter.HandleFunc("/field/{collectionname}", UpdateFieldName).Methods("POST")

	myRouter.HandleFunc("/addfield/{collectionname}", AddField).Methods("POST")
	myRouter.HandleFunc("/addcollection", AddNewCollection).Methods("POST")

	myRouter.HandleFunc("/fields/{collectionname}", UpdateFieldNames).Methods("POST")
	myRouter.HandleFunc("/meta/{collectionname}", ReturnCollectionMeta).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(corsObj, headersOk, methodsOk)(myRouter)))
}

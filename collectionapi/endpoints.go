package collectionapi

import (
	"encoding/json"
	"fmt"
	coll "gencollection/collectionstructure"
	"net/http"

	"github.com/gorilla/mux"
)

// ReturnCollectionAsTable writes the collection
// data in a table-format.
func ReturnCollectionAsTable(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: ReturnCollectionAsTable")
	vars := mux.Vars(r)
	key := vars["collectionname"]
	f := coll.Read("collectionsdata.json")
	collection := f.PullCollection(key)
	json.NewEncoder(w).Encode(collection.RecordsToTable())

}

// CollectionNames obtains the meta data associated
// with all collections. Result written as json
// to the calling http request.
func CollectionNames(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: CollectionNames")
	f := coll.Read("collectionsdata.json")
	json.NewEncoder(w).Encode(f.Names())
}

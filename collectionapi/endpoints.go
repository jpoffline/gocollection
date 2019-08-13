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
	f := coll.Read(coll.DATAFILE)
	collection := f.PullCollection(key)
	json.NewEncoder(w).Encode(collection.RecordsToTable())

}

// CollectionNames obtains the meta data associated
// with all collections. Result written as json
// to the calling http request.
func CollectionNames(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: CollectionNames")
	json.NewEncoder(w).Encode(coll.GetMeta())
}

func AddRecordToCollection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: AddRecordToCollection")
	w.Header().Set("Content-Type", "application/json")

	var record coll.RecordReceive
	vars := mux.Vars(r)
	key := vars["collectionname"]
	_ = json.NewDecoder(r.Body).Decode(&record)
	coll.AddRecord(key, record)

	json.NewEncoder(w).Encode(&record)
}

func ReturnFieldNames(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: ReturnFieldNames")
	vars := mux.Vars(r)
	key := vars["collectionname"]
	json.NewEncoder(w).Encode(coll.GetFields(key))
}

func UpdateFieldName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: UpdateFieldName")
	vars := mux.Vars(r)
	key := vars["collectionname"]
	var field coll.Field
	_ = json.NewDecoder(r.Body).Decode(&field)

	coll.UpdateFieldName(key, field)
	//json.NewEncoder(w).Encode()
}

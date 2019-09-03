package collectionapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	coll "github.com/jpoffline/gocollection/collectionstructure"

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

// AddRecordToCollection is an endpoint which adds
// a given record to a collection.

func AddRecordToCollection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: AddRecordToCollection")
	w.Header().Set("Content-Type", "application/json")

	var record coll.RecordReceive
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
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
}

func UpdateFieldNames(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: UpdateFieldNames")
	vars := mux.Vars(r)
	key := vars["collectionname"]
	var fields []coll.Field

	_ = json.NewDecoder(r.Body).Decode(&fields)
	for _, field := range fields {
		coll.UpdateFieldName(key, field)
	}
}

func ReturnCollectionMeta(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: ReturnCollectionMeta")
	vars := mux.Vars(r)
	key := vars["collectionname"]

	json.NewEncoder(w).Encode(coll.GetCollectionMeta(key))

}

func AddField(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: AddField")
	vars := mux.Vars(r)
	key := vars["collectionname"]

	w.Header().Set("Content-Type", "application/json")

	var fn map[string]string

	w.WriteHeader(http.StatusOK)

	_ = json.NewDecoder(r.Body).Decode(&fn)
	coll.AddField(key, fn["newfield"])

	json.NewEncoder(w).Encode(&fn)
}

package collectionstructure

const DATAFILE = "collectionsdata.json"

// Collections is the main collections data structure.
type Collections struct {
	CollectionNames []CollectionMeta      `json:"collection_names"`
	Data            map[string]Collection `json:"collections"`
}

// Collection contains the fields and records.
type Collection struct {
	Fields  []Field  `json:"fields"`
	Records []Record `json:"records"`
}

// Field contains the meta data for a given field.
type Field struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Record contains the data for a particular record.
type Record struct {
	ID   int               `json:"id"`
	Data map[string]string `json:"data"`
}

type RecordReceive struct {
	ID   int               `json:"id"`
	Data map[string]string `json:"data"`
}

// CollectionMeta contains the meta data for a collection.
type CollectionMeta struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

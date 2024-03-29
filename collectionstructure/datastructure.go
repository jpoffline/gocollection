package collectionstructure

const DATAFILE = "collectionsdata.json"

// Collections is the main collections data structure.
type Collections struct {
	CollectionNames []CollectionMeta      `json:"collection_names"`
	Data            map[string]Collection `json:"collections"`
}

// Collection contains the fields and records.
type Collection struct {
	MaxFieldIdx int      `json:"maxfieldidx"`
	Fields      []Field  `json:"fields"`
	Records     []Record `json:"records"`
}

// Field contains the meta data for a given field.
type Field struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FieldsReturn struct {
	CollectionName string  `json:"collection"`
	Fields         []Field `json:"fields"`
}

// Record contains the data for a particular record.
type Record struct {
	ID   int        `json:"id"`
	Data RecordItem `json:"data"`
}

//RecordReceive is the struct of a record
// recived.
type RecordReceive struct {
	//ID   int               `json:"id"`
	Data RecordItem `json:"data"`
}

// CollectionMeta contains the meta data for a collection.
type CollectionMeta struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

// RecordItem is the type of a record item.
type RecordItem map[string]string

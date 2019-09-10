package collectionstructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Names returns the list of collection meta data.
func (cols *Collections) Names() []CollectionMeta {
	return (cols.CollectionNames)
}

func Read(filename string) Collections {
	fmt.Println(filename)
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file%q", err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error reading file%q", err)
	}
	var f Collections
	json.Unmarshal(byteValue, &f)
	return (f)
}

// NumRecords returns the number of records in a given
// collection.
func (coll *Collection) NumRecords() int {
	return len(coll.Records)
}

// Write will write the collections to the provided filename.
func (coll *Collections) Write(filename string) {
	file, _ := json.MarshalIndent(coll, "", " ")
	_ = ioutil.WriteFile(filename, file, 0644)
}

// Save a given collection.
func (coll Collection) Save(colname string) {
	current := Read(DATAFILE)
	current.Data[colname] = coll
	current.Write(DATAFILE)
}

func CollectionToTable(colname string) Table {
	f := Read(DATAFILE)
	collection := f.PullCollection(colname)
	return (collection.RecordsToTable())
}

// GetFields will return the fields of a provided collection.
func GetFields(colname string) FieldsReturn {
	current := Read(DATAFILE)
	coll := current.PullCollection(colname)
	var ret FieldsReturn
	ret.Fields = coll.pullFields()
	ret.CollectionName = colname
	return (ret)
}

// GetMeta will return the names of all collections.
func GetMeta() []CollectionMeta {
	f := Read(DATAFILE)
	return (f.Names())
}

// GetCollectionMeta will return the meta data for the provided collection name.
func GetCollectionMeta(collname string) CollectionMeta {
	all := GetMeta()
	for _, v := range all {
		if v.Name == collname {
			return v
		}
	}
	return CollectionMeta{}
}

// AddRecord will add a record to the provided collection name.
func AddRecord(colname string, rec RecordReceive) {

	current := Read(DATAFILE)
	tgtCollection := current.PullCollection(colname)

	tgtCollection.addRecord(rec)

	tgtCollection.Save(colname)
}

// UpdateFieldName will update the field name of a give collection.
func UpdateFieldName(colname string, field Field) {
	current := Read(DATAFILE)
	coll := current.PullCollection(colname)
	coll.modifyFieldName(field)
	coll.Save(colname)
}

// UpdateFieldNames will update the field names of the provided
// collection name.
func UpdateFieldNames(colname string, fields []Field) {
	for _, field := range fields {
		UpdateFieldName(colname, field)
	}
}

// AddField will add a fieldname to the provided collection name.
func AddField(colname, fieldname string) {
	current := Read(DATAFILE)
	coll := current.PullCollection(colname)
	coll.addField(fieldname)
	coll.Save(colname)
}

// AddCollection adds a collection with the provided meta.
func AddCollection(meta CollectionMeta) {
	current := Read(DATAFILE)

	current.addNewCollection(meta)

	current.Write(DATAFILE)
}

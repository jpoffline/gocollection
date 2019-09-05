package collectionstructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

func (coll *Collections) Write(filename string) {
	file, _ := json.MarshalIndent(coll, "", " ")
	_ = ioutil.WriteFile(filename, file, 0644)
}

func saveCollection(colname string, col Collection) {
	current := Read(DATAFILE)
	current.Data[colname] = col
	current.Write(DATAFILE)
}

func (coll *Collections) PullCollection(colname string) Collection {
	return (coll.Data[colname])
}

func (coll *Collection) pullFields() []Field {
	return (coll.Fields)
}

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

func (coll *Collection) addRecord(r Record) {
	r.ID = len(coll.Records) + 1
	coll.Records = append(coll.Records, r)
}

// AddRecord will add a record to the provided collection name.
func AddRecord(colname string, rec RecordReceive) {

	current := Read(DATAFILE)
	tgtCollection := current.PullCollection(colname)
	tgtFields := tgtCollection.pullFields()

	recData := newRecordItem()
	for _, v := range tgtFields {
		ii, ok := rec.Data[v.Name]

		if ok {
			recData[v.ID] = ii
		}

	}

	var newRecord Record
	newRecord.Data = recData
	tgtCollection.addRecord(newRecord)
	saveCollection(colname, tgtCollection)
}

func UpdateFieldName(colname string, field Field) {
	current := Read(DATAFILE)
	coll := current.PullCollection(colname)
	coll.modifyFieldName(field)
	saveCollection(colname, coll)
}

func (coll *Collection) modifyFieldName(field Field) {
	fields := coll.pullFields()

	for i, f := range fields {
		if f.ID == field.ID {
			fields[i].Name = field.Name
		}
	}
	coll.Fields = fields
}

func (coll *Collection) addField(fieldname string) {
	fields := coll.pullFields()

	newMax := coll.MaxFieldIdx + 1
	var newfield Field
	newfield.ID = strconv.Itoa(newMax)
	newfield.Name = fieldname
	fields = append(fields, newfield)
	coll.Fields = fields
	coll.MaxFieldIdx = newMax
}

// AddField will add a fieldname to the provided collection name.
func AddField(colname, fieldname string) {
	current := Read(DATAFILE)
	coll := current.PullCollection(colname)
	coll.addField(fieldname)
	saveCollection(colname, coll)
}

func (cols *Collections) addNewCollection(meta CollectionMeta) {
	currCollMeta := cols.CollectionNames
	currMaxID, _ := strconv.Atoi(currCollMeta[len(currCollMeta)-1].ID)
	meta.ID = strconv.Itoa(currMaxID + 1)
	cols.CollectionNames = append(currCollMeta, meta)
	cols.Data[meta.Name] = newCollectionStructure()
}

// AddCollection adds a collection with the provided meta.
func AddCollection(meta CollectionMeta) {
	current := Read(DATAFILE)

	current.addNewCollection(meta)

	current.Write(DATAFILE)
}

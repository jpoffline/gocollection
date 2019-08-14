package collectionstructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Names returns the list of collection meta data.
func (col *Collections) Names() []CollectionMeta {
	return (col.CollectionNames)
}

func Read(filename string) Collections {
	jsonFile, _ := os.Open(filename)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var f Collections
	json.Unmarshal(byteValue, &f)
	return (f)
}

func (col *Collections) Write(filename string) {
	file, _ := json.MarshalIndent(col, "", " ")
	_ = ioutil.WriteFile(filename, file, 0644)
}

func saveCollection(colname string, col Collection) {
	current := Read(DATAFILE)
	current.Data[colname] = col
	current.Write(DATAFILE)
}

func (col *Collections) PullCollection(colname string) Collection {
	return (col.Data[colname])
}

func (c *Collection) pullFields() []Field {
	return (c.Fields)
}

func GetFields(colname string) FieldsReturn {
	current := Read(DATAFILE)
	coll := current.PullCollection(colname)
	var ret FieldsReturn
	ret.Fields = coll.pullFields()
	ret.CollectionName = colname
	return (ret)
}

func GetMeta() []CollectionMeta {
	f := Read(DATAFILE)
	return (f.Names())
}

func GetCollectionMeta(collname string) CollectionMeta {
	all := GetMeta()
	for _, v := range all {
		if v.Name == collname {
			return v
		}
	}
	return CollectionMeta{}
}

func (c *Collection) addRecord(r Record) {
	c.Records = append(c.Records, r)
}

func AddRecord(colname string, rec RecordReceive) {
	fmt.Println(rec)
	current := Read(DATAFILE)
	tgtCollection := current.PullCollection(colname)
	tgtFields := tgtCollection.pullFields()

	var newRecord Record
	newRecord.ID = rec.ID

	recData := make(map[string]string)
	for _, v := range tgtFields {
		ii, ok := rec.Data[v.Name]

		if ok {
			recData[v.ID] = ii
		}

	}

	newRecord.Data = recData
	tgtCollection.addRecord(newRecord)
	saveCollection(colname, tgtCollection)
}

func UpdateFieldName(colname string, field Field) {
	current := Read(DATAFILE)
	coll := current.PullCollection(colname)
	fields := coll.pullFields()
	for i, f := range fields {
		if f.ID == field.ID {
			fields[i].Name = field.Name
		}
	}
	coll.Fields = fields
	saveCollection(colname, coll)
}

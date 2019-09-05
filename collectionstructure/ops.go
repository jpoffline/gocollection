package collectionstructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Names returns the list of collection meta data.
func (col *Collections) Names() []CollectionMeta {
	return (col.CollectionNames)
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
func (col *Collection) NumRecords() int {
	return len(col.Records)
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

func (col *Collection) addRecord(r Record) {
	r.ID = len(col.Records) + 1
	col.Records = append(col.Records, r)
}

func AddRecord(colname string, rec RecordReceive) {

	current := Read(DATAFILE)
	tgtCollection := current.PullCollection(colname)
	tgtFields := tgtCollection.pullFields()

	var newRecord Record
	//newRecord.ID = rec.ID

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
	//
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

func AddField(colname, fieldname string) {
	current := Read(DATAFILE)
	coll := current.PullCollection(colname)
	fields := coll.pullFields()

	newMax := coll.MaxFieldIdx + 1
	var newfield Field
	newfield.ID = strconv.Itoa(newMax)
	newfield.Name = fieldname
	fields = append(fields, newfield)
	coll.Fields = fields
	coll.MaxFieldIdx = newMax

	saveCollection(colname, coll)
}

func newCollectionStructure() Collection {
	newColl := Collection{}
	newColl.MaxFieldIdx = 1
	var nf Field

	nf.ID = "1"
	nf.Name = "Field"
	newColl.Fields = append(newColl.Fields, nf)
	var dd = make(map[string]string)
	dd[nf.ID] = "Record"
	var nr Record
	nr.ID = 1
	nr.Data = dd
	newColl.Records = append(newColl.Records, nr)
	return newColl
}

func AddCollection(meta CollectionMeta) {
	current := Read(DATAFILE)
	currCollMeta := current.CollectionNames
	currMaxID, _ := strconv.Atoi(currCollMeta[len(currCollMeta)-1].ID)
	meta.ID = strconv.Itoa(currMaxID + 1)
	current.CollectionNames = append(currCollMeta, meta)

	current.Data[meta.Name] = newCollectionStructure()

	current.Write(DATAFILE)
}

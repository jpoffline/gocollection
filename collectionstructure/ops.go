package collectionstructure

import (
	"encoding/json"
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

func (col *Collections) PullCollection(colname string) Collection {
	return (col.Data[colname])
}

func (c *Collection) pullFields() []Field {
	return (c.Fields)
}

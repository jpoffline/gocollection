package collectionstructure

import "strconv"

func (cols *Collections) addNewCollection(meta CollectionMeta) {
	currCollMeta := cols.CollectionNames
	currMaxID, _ := strconv.Atoi(currCollMeta[len(currCollMeta)-1].ID)
	meta.ID = strconv.Itoa(currMaxID + 1)
	cols.CollectionNames = append(currCollMeta, meta)
	cols.Data[meta.Name] = newCollectionStructure()
}

// PullCollection will extract a particular collection from the collections.
func (cols *Collections) PullCollection(colname string) Collection {
	return (cols.Data[colname])
}

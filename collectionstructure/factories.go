package collectionstructure

func newField(id, name string) Field {
	return Field{ID: id, Name: name}
}

func newCollectionStructure() Collection {
	newColl := Collection{}
	newColl.MaxFieldIdx = 1
	nf := newField("1", "Field")

	newColl.Fields = append(newColl.Fields, nf)
	var dd = newRecordItem()
	dd[nf.ID] = "Record"
	var nr Record
	nr.ID = 1
	nr.Data = dd
	newColl.Records = append(newColl.Records, nr)
	return newColl
}

func newRecordItem() RecordItem {
	return (make(RecordItem))
}

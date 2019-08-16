package collectionstructure

import "testing"

func TestAddRecord(t *testing.T) {

	var r = Record{}

	coll := Collection{}

	if coll.NumRecords() != 0 {
		t.Errorf("NOT")
	}

	coll.addRecord(r)

	if coll.NumRecords() != 1 {
		t.Errorf("NOT")
	}

	if coll.Records[0].ID != 1 {
		t.Error("NOT RIGHT ID")
	}

}

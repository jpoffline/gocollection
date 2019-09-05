package collectionstructure

import "strconv"

func (coll *Collection) pullFields() []Field {
	return (coll.Fields)
}

func (coll *Collection) addRecord(rec RecordReceive) {
	recData := coll.prepNewRecord(rec)
	coll.pushRecord(recData)
}

func (coll *Collection) pushRecord(ri RecordItem) {
	var r Record
	r.Data = ri
	r.ID = len(coll.Records) + 1
	coll.Records = append(coll.Records, r)
}

func (coll *Collection) prepNewRecord(rec RecordReceive) RecordItem {
	tgtFields := coll.pullFields()

	recData := newRecordItem()
	for _, v := range tgtFields {
		ii, ok := rec.Data[v.Name]

		if ok {
			recData[v.ID] = ii
		}

	}

	return recData
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

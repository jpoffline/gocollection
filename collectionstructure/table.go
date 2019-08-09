package collectionstructure

import "strconv"

type Table struct {
	Header []string
	Rows   [][]string
}

func (c *Collection) RecordsToTable() Table {
	var ids []string
	var tbl Table
	tbl.Header = append(tbl.Header, "ID")
	for _, f := range c.Fields {
		ids = append(ids, f.ID)
		tbl.Header = append(tbl.Header, f.Name)
	}
	for _, r := range c.Records {
		iid := strconv.Itoa(r.ID)
		row := []string{iid}
		for _, id := range ids {
			ii, ok := r.Data[id]
			var item string
			if ok {
				item = ii.(string)
			} else {
				item = "NONE"
			}

			row = append(row, item)
		}
		tbl.Rows = append(tbl.Rows, row)

	}
	return (tbl)
}

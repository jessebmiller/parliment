package format

import (
	"fmt"
	"database/sql"
	"encoding/json"
	"log"
)

type Formatter interface {
	Format() []byte
}

type SimpleJsonResult struct {
	rows sql.Rows
}

func NewSimpleJsonResult(rows sql.Rows) SimpleJsonResult {
	return SimpleJsonResult{rows}
}

func (r SimpleJsonResult) Format() []byte {
	rows := make([][]string, 0)
	columns, err := r.rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	for r.rows.Next() {
		rawRow := make([]interface{}, len(columns))
		r.rows.Scan(rawRow...)
		row := make([]string, len(rawRow))
		for elem, _ := range rawRow {
			row = append(row, fmt.Sprintf("%v", elem))
		}
		rows = append(rows, row)
	}

	columnTypes, err := r.rows.ColumnTypes()
	if err != nil {
		log.Fatal(err)
	}
	types := make([]string, 0)
	for _, columnType := range columnTypes {
		types = append(types, columnType.DatabaseTypeName())
	}

	result := struct{
		columns []string
		types []string
		rows [][]string
	}{
		columns,
		types,
		rows,
	}

	formattedResult, err := json.Marshal(result)
	if err != nil {
		// TODO handle this gracefully
		log.Fatal(err)
	}
	return formattedResult
}

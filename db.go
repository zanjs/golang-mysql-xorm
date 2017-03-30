package main

import (
	"database/sql"
)

// Record is ...
func Record(rows *sql.Rows) (record map[string]string) {

	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))

	values := make([]interface{}, len(columns))

	for j := range values {

		scanArgs[j] = &values[j]

	}

	record = make(map[string]string)

	for rows.Next() {

		//将行数据保存到record字典

		err := rows.Scan(scanArgs...)

		CheckErr(err)

		for i, col := range values {

			if col != nil {

				record[columns[i]] = string(col.([]byte))

			}

		}

	}

	return
}

// RSQL is ...
func RSQL(sql string) (record map[string]string) {
	// rows, _ := DB.Query("SELECT * FROM user")
	rows, err := DB.Query(sql)
	defer rows.Close()

	record = Record(rows)

	CheckErr(err)

	return record
}

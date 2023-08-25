package db

import "database/sql"

var DB *sql.DB

type dbObject interface {
	scan(*sql.Rows) (dbObject, error)
}

func query[T dbObject](query string, args ...any) ([]T, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	var objects []T
	var objectScan T
	for rows.Next() {
		object, err := objectScan.scan(rows)
		if err != nil {
			return objects, err
		}
		objects = append(objects, object.(T))
	}
	return objects, nil
}

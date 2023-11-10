package db

import (
	"database/sql"

	"github.com/t0k4r/qb"
)

var DB *sql.DB

func Query[T qb.Selectable](query *qb.QSelect, args ...any) ([]T, error) {
	return qb.Query[T](query, DB, args...)
}

func QueryFirst[T qb.Selectable](query *qb.QSelect, args ...any) (*T, error) {
	items, err := qb.Query[T](query, DB, args...)
	if len(items) == 0 {
		return nil, err
	}
	return &items[0], err
}

package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/murphybytes/journal/server/journal"
)

type database struct {
	db *sqlx.DB
}

// New creates database struct to access mysql
func New(dsn string) (journal.Database, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	ds := &database{
		db: db,
	}
	return ds, nil
}

func (d *database) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}

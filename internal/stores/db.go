package stores

import (
	"database/sql"
	"tgb/internal/stores/sql_start"
)
import _ "github.com/mattn/go-sqlite3"

type Store struct {
	db *sql.DB
}

func New() (*Store, error) {
	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		return nil, err
	}

	if _, err = db.Exec(sql_start.CreateTable); err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

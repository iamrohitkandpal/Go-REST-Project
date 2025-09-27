package sqlite

import (
	"database/sql"

	"github.com/iamrohitkandpal/Go-REST-Project/internal/config"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
}
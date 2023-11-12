package database

import (
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
)

type Seed struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Seed {
	return Seed{
		db: db,
	}
}

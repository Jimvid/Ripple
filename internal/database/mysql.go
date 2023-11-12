package database

import (
	"github.com/jimvid/ripple/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func CreateMySqlConnection(env config.EnvVars) *sqlx.DB {

	db := sqlx.MustConnect("mysql", env.DB_CONNECTION_STRING)

	err := db.Ping()
	if err != nil {
		panic(err)
	} else {
		println("DB CONNECTED")
	}

	return db
}

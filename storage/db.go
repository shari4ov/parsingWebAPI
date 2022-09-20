package storage

import (
	config "api/config"
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConnection() *sql.DB {
	db, err := sql.Open("postgres", config.GetPostgresConnection())
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

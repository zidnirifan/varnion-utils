package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Ping(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Ping()
}

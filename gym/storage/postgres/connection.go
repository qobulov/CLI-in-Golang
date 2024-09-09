package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres",
		"host=localhost port=5432 user=postgres password=BEKJONS dbname=gym sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}

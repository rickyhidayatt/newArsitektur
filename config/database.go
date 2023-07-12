package config

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB
var singleton sync.Once

func DB() *sql.DB {
	singleton.Do(func() {
		var err error
		if db == nil {
			dbConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				DATABASE_HOST,
				DATABASE_PORT,
				DATABASE_USER,
				DATABASE_PASS,
				DATABASE_NAME)

			db, err = sql.Open("postgres", dbConnection)
			if err != nil {
				panic(err)
			}
		}
	})

	return db
}

package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	config "bni.co.id/xpora/medias/config"
	_ "github.com/lib/pq"
)

var db *sql.DB
var singleton sync.Once

func DB() *sql.DB {
	singleton.Do(func() {
		if db == nil {
			var err error
			connStr := fmt.Sprintf(
				"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
				config.GetEnv(config.DATABASE_USER),
				config.GetEnv(config.DATABASE_PASS),
				config.GetEnv(config.DATABASE_HOST),
				config.GetEnv(config.DATABASE_PORT),
				config.GetEnv(config.DATABASE_NAME),
				config.GetEnv(config.DATABASE_SSL),
			)

			db, err = sql.Open("postgres", connStr)

			if err != nil {
				log.Fatal(err)
			}
		}
	})

	return db
}

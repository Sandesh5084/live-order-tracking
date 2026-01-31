package postgres

import (
	"database/sql"
	"log"
	"time"
	"mylotapp/internal/config"

	_ "github.com/lib/pq"
)

func MustConnectDB(cfg config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatal("database not reachable:", err)
	}

	return db
}

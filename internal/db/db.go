package db

import (
	"algoBot/internal/config"
	"database/sql"
	"log"
)

func InitDB(cfg config.Config) *sql.DB {
	db, err := sql.Open(cfg.DriverName, cfg.ConnString)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}

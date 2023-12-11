package db

import (
	"algoBot/internal/config"
	"database/sql"
	"log/slog"
)

func InitDB(cfg *config.Config, log *slog.Logger) *sql.DB {
	db, err := sql.Open(cfg.DriverName, cfg.ConnString)
	if err != nil {
		log.Error("can't open db", slog.String("err", err.Error()))
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Error("can't ping db", slog.String("err", err.Error()))
		return nil
	}
	return db
}

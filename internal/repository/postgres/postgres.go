package postgres

import (
	"algoBot/internal/repository"
	"context"
	"database/sql"
	"log"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Get(ctx context.Context, topic string) (*repository.Page, error) {
	query := `
	SELECT topic, short_desc, useful_materials
	FROM algo_and_structures
	WHERE topic = $1`

	rows, err := r.db.QueryContext(ctx, query, topic)
	if err != nil {
		log.Printf("query(get) execution error: %s", err.Error())
		return nil, err
	}
	var shortDesc, usefulMaterials string
	for rows.Next() {
		err = rows.Scan(&topic, &shortDesc, &usefulMaterials)
		if err != nil {
			log.Printf("can't scan rows: %s", err.Error())
			return nil, err
		}
	}
	if rows.Err() != nil {
		log.Printf("rows error: %s", rows.Err())
		return nil, rows.Err()
	}

	return &repository.Page{
		Topic:           topic,
		ShortDesc:       shortDesc,
		UsefulMaterials: usefulMaterials,
	}, nil
}

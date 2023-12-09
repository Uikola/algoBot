package repository

import "context"

type Page struct {
	Topic           string
	ShortDesc       string
	UsefulMaterials string
}

type Repository interface {
	Get(ctx context.Context, topic string) (*Page, error)
}

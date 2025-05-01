package categories_repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoriesRepository struct {
	db *pgxpool.Pool
}

func NewCategoriesRepository(db *pgxpool.Pool) *CategoriesRepository {
	return &CategoriesRepository{
		db: db,
	}
}

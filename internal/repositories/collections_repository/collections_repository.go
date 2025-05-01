package collections_repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	CollectionRepository struct {
		pool *pgxpool.Pool
	}
)

// var _ charity_service.IColleсtionsRepository = (*CollectionRepository)(nil)

func NewCollectionRepository(pool *pgxpool.Pool) *CollectionRepository {
	return &CollectionRepository{
		pool: pool,
	}
}

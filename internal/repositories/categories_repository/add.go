package categories_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/charity/internal/services"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

func (r *CategoriesRepository) Add(ctx context.Context, category string) (int64, error) {
	if category == "" {
		return -1, services.InvalidName
	}

	var id int64
	query := `INSERT INTO categories (name) VALUES ($1) RETURNING id`
	err := r.db.QueryRow(ctx, query, category).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return -1, services.AlreadyExists
		}

		log.Error().Msg(err.Error())
		return -1, services.InternalServerError
	}
	return id, nil
}

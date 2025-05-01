package categories_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/charity/internal/models"
	"github.com/GP-Hacks/charity/internal/services"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func (r *CategoriesRepository) GetAll(ctx context.Context) ([]*models.Category, error) {
	query := `SELECT id, name FROM categories`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, services.InternalServerError
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name); err != nil {
			log.Error().Msg(err.Error())
			return nil, services.InternalServerError
		}
		categories = append(categories, &cat)
	}

	if err := rows.Err(); err != nil {
		log.Error().Msg(err.Error())
		return nil, services.InternalServerError
	}

	return categories, nil
}

func (r *CategoriesRepository) GetById(ctx context.Context, id int64) (*models.Category, error) {
	query := `SELECT id, name FROM categories WHERE id = $1`

	var cat models.Category
	err := r.db.QueryRow(ctx, query, id).Scan(
		&cat.ID,
		&cat.Name,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, services.NotFound
		}

		log.Error().Msg(err.Error())
		return nil, services.InternalServerError
	}

	return &cat, nil
}

func (r *CategoriesRepository) GetByName(ctx context.Context, name string) (*models.Category, error) {
	query := `SELECT id, name FROM categories WHERE name = $1`

	var cat models.Category
	err := r.db.QueryRow(ctx, query, name).Scan(
		&cat.ID,
		&cat.Name,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, services.NotFound
		}

		log.Error().Msg(err.Error())
		return nil, services.InternalServerError
	}

	return &cat, nil
}

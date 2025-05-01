package collections_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/charity/internal/models"
	"github.com/GP-Hacks/charity/internal/services"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func (cr *CollectionRepository) Get(ctx context.Context, offset, limit int64) ([]*models.Collection, error) {
	rows, err := cr.pool.Query(ctx, `
		SELECT
			c.id,
			c.category_id,
			c.name,
			c.description,
			c.orgranization,
			c.phone,
			c.website,
			c.goal,
			c.current,
			c.photo
		FROM collections c
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, services.InternalServerError
	}
	defer rows.Close()

	var colls []*models.Collection
	for rows.Next() {
		var coll models.Collection
		err := rows.Scan(
			&coll.ID,
			&coll.CategoryID,
			&coll.Name,
			&coll.Description,
			&coll.Organization,
			&coll.Phone,
			&coll.Website,
			&coll.Goal,
			&coll.Current,
			&coll.Photo,
		)
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, services.InternalServerError
		}

		colls = append(colls, &coll)
	}

	return colls, nil
}

func (cr *CollectionRepository) GetByCategory(ctx context.Context, offset, limit, categoryId int64) ([]*models.Collection, error) {
	rows, err := cr.pool.Query(ctx, `
		SELECT
			c.id,
			c.category_id,
			c.name,
			c.description,
			c.orgranization,
			c.phone,
			c.website,
			c.goal,
			c.current,
			c.photo
		FROM collections c
		WHERE c.category_id == $1
		LIMIT $2 OFFSET $3
	`, categoryId, limit, offset)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, services.InternalServerError
	}
	defer rows.Close()

	var colls []*models.Collection
	for rows.Next() {
		var coll models.Collection
		err := rows.Scan(
			&coll.ID,
			&coll.CategoryID,
			&coll.Name,
			&coll.Description,
			&coll.Organization,
			&coll.Phone,
			&coll.Website,
			&coll.Goal,
			&coll.Current,
			&coll.Photo,
		)
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, services.InternalServerError
		}

		colls = append(colls, &coll)
	}

	return colls, nil
}

func (cr *CollectionRepository) GetById(ctx context.Context, id int64) (*models.Collection, error) {
	query := `SELECT id, name, category_id, description, orgranization, phone, website, goal, current, photo FROM categories WHERE id = $1`

	var coll models.Collection
	err := cr.pool.QueryRow(ctx, query, id).Scan(
		&coll.ID,
		&coll.Name,
		&coll.CategoryID,
		&coll.Description,
		&coll.Organization,
		&coll.Phone,
		&coll.Website,
		&coll.Goal,
		&coll.Current,
		&coll.Photo,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, services.NotFound
		}

		log.Error().Msg(err.Error())
		return nil, services.InternalServerError
	}

	return &coll, nil
}

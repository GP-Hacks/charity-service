package collections_repository

import (
	"context"

	"github.com/GP-Hacks/charity/internal/models"
	"github.com/GP-Hacks/charity/internal/services"
	"github.com/rs/zerolog/log"
)

func (cr *CollectionRepository) Add(ctx context.Context, coll *models.Collection) (int64, error) {
	query := `
		INSERT INTO collections (
			category_id,
			name,
			description,
			organization,
			phone,
			website,
			goal,
			current,
			photo
		) VALUES ($1, $2, $3, $4, $5, $6, $7, 0, $8)
		RETURNING id
	`

	var id int64
	err := cr.pool.QueryRow(ctx, query,
		coll.CategoryID,
		coll.Name,
		coll.Description,
		coll.Organization,
		coll.Phone,
		coll.Website,
		coll.Goal,
		coll.Photo,
	).Scan(&id)

	if err != nil {
		log.Error().Msg(err.Error())
		return 0, services.InternalServerError
	}

	return id, nil
}

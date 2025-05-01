package collections_repository

import (
	"context"

	"github.com/GP-Hacks/charity/internal/models"
	"github.com/GP-Hacks/charity/internal/services"
	"github.com/rs/zerolog/log"
)

func (cr *CollectionRepository) Update(ctx context.Context, coll *models.Collection) error {
	query := `
		UPDATE collections
		SET 
			category_id = $1,
			name = $2,
			description = $3,
			organization = $4,
			phone = $5,
			website = $6,
			goal = $7,
			photo = $8
		WHERE id = $9
	`

	_, err := cr.pool.Exec(ctx, query,
		coll.CategoryID,
		coll.Name,
		coll.Description,
		coll.Organization,
		coll.Phone,
		coll.Website,
		coll.Goal,
		coll.Photo,
		coll.ID,
	)
	if err != nil {
		log.Error().Msg(err.Error())
		return services.InternalServerError
	}
	return nil
}

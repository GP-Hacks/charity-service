package charity_service

import (
	"context"

	"github.com/GP-Hacks/charity/internal/models"
	"github.com/GP-Hacks/charity/internal/services"
)

func (cs *CharityService) AddCollection(ctx context.Context, collection *models.Collection, token string) error {
	user, err := cs.userAdapter.GetByToken(ctx, token)
	if err != nil {
		return err
	}

	if user.Status != models.Admin {
		return services.AccessDenied
	}

	_, err = cs.collectionsRepository.Add(ctx, collection)
	return err
}

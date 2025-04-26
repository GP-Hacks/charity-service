package charity_service

import (
	"context"

	"github.com/GP-Hacks/kdt2024-charity/internal/models"
	"github.com/GP-Hacks/kdt2024-charity/internal/services"
)

func (cs *CharityService) AddCollection(ctx context.Context, collection *models.Collection, token string) error {
	user, err := cs.userAdapter.GetUserByToken(ctx, token)
	if err != nil {
		return err
	}

	if user.Status != models.Admin {
		return services.AccessDenied
	}

	return cs.collectionsRepository.AddCollection(ctx, collection)
}

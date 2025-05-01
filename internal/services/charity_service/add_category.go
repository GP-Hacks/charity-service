package charity_service

import (
	"context"

	"github.com/GP-Hacks/charity/internal/models"
	"github.com/GP-Hacks/charity/internal/services"
)

func (cs *CharityService) AddCategory(ctx context.Context, category, token string) error {
	user, err := cs.userAdapter.GetByToken(ctx, token)
	if err != nil {
		return err
	}

	if user.Status != models.Admin {
		return services.AccessDenied
	}

	_, err = cs.categoriesRepository.Add(ctx, category)
	return err
}

package charity_service

import (
	"context"
	"slices"

	"github.com/GP-Hacks/kdt2024-charity/internal/models"
	"github.com/GP-Hacks/kdt2024-charity/internal/services"
)

func (cs *CharityService) GetCollections(ctx context.Context, offset, limit int64) ([]*models.Collection, error) {
	return cs.collectionsRepository.GetCollections(ctx, offset, limit)
}

func (cs *CharityService) GetCollectionsByCategory(ctx context.Context, category string, offset, limit int64) ([]*models.Collection, error) {
	categories, err := cs.categoriesRepository.Get(ctx)
	if err != nil {
		return nil, err
	}

	if !slices.Contains(categories, category) {
		return nil, services.CategoryNotFound
	}

	return cs.collectionsRepository.GetCollectionsByCategory(ctx, category, offset, limit)
}

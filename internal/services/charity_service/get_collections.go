package charity_service

import (
	"context"

	"github.com/GP-Hacks/charity/internal/models"
)

func (cs *CharityService) GetCollections(ctx context.Context, offset, limit int64) ([]*models.CollectionWithCategory, error) {
	colls, err := cs.collectionsRepository.Get(ctx, offset, limit)
	if err != nil {
		return []*models.CollectionWithCategory{}, err
	}

	res := make([]*models.CollectionWithCategory, len(colls))
	for i, coll := range colls {
		cat, err := cs.categoriesRepository.GetById(ctx, coll.ID)
		if err != nil {
			return []*models.CollectionWithCategory{}, err
		}

		res[i] = coll.ToCollectionWithCategory(cat.Name)
	}

	return res, nil
}

func (cs *CharityService) GetCollectionsByCategory(ctx context.Context, category string, offset, limit int64) ([]*models.CollectionWithCategory, error) {
	cat, err := cs.categoriesRepository.GetByName(ctx, category)
	if err != nil {
		return []*models.CollectionWithCategory{}, err
	}

	colls, err := cs.collectionsRepository.GetByCategory(ctx, offset, limit, cat.ID)
	if err != nil {
		return []*models.CollectionWithCategory{}, err
	}

	res := make([]*models.CollectionWithCategory, len(colls))
	for i, coll := range colls {
		res[i] = coll.ToCollectionWithCategory(category)
	}

	return res, nil

}

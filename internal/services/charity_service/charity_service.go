package charity_service

import (
	"context"

	"github.com/GP-Hacks/kdt2024-charity/internal/models"
)

type (
	ICategoriesRepository interface {
		Get(ctx context.Context) ([]string, error)
		Add(ctx context.Context, category string) error
	}

	IColleсtionsRepository interface {
		GetCollections(ctx context.Context, offset, limit int64) ([]*models.Collection, error)
		GetCollectionsByCategory(ctx context.Context, category string, offset, limit int64) ([]*models.Collection, error)
		UpdateCollection(ctx context.Context, collection *models.Collection) error
		AddCollection(ctx context.Context, collection *models.Collection) error
		GetById(ctx context.Context, id int64) (*models.Collection, error)
	}

	IUserAdapter interface {
		GetUserByToken(ctx context.Context, token string) (*models.User, error)
	}

	CharityService struct {
		categoriesRepository  ICategoriesRepository
		collectionsRepository IColleсtionsRepository
		userAdapter           IUserAdapter
	}
)

func NewCharityService(categoriesRepo ICategoriesRepository, collectionsRepo IColleсtionsRepository, userAdapter IUserAdapter) *CharityService {
	return &CharityService{
		categoriesRepository:  categoriesRepo,
		collectionsRepository: collectionsRepo,
		userAdapter:           userAdapter,
	}
}

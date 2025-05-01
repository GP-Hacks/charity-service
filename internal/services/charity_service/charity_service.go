package charity_service

import (
	"context"

	"github.com/GP-Hacks/charity/internal/models"
)

type (
	ICategoriesRepository interface {
		GetAll(ctx context.Context) ([]*models.Category, error)
		GetById(ctx context.Context, id int64) (*models.Category, error)
		GetByName(ctx context.Context, name string) (*models.Category, error)
		Add(ctx context.Context, category string) (int64, error)
	}

	IColleсtionsRepository interface {
		Add(ctx context.Context, coll *models.Collection) (int64, error)
		Get(ctx context.Context, offset, limit int64) ([]*models.Collection, error)
		GetByCategory(ctx context.Context, offset, limit, categoryId int64) ([]*models.Collection, error)
		GetById(ctx context.Context, id int64) (*models.Collection, error)
		Update(ctx context.Context, coll *models.Collection) error
	}

	IUserAdapter interface {
		GetByToken(ctx context.Context, token string) (*models.User, error)
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

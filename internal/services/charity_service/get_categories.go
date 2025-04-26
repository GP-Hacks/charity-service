package charity_service

import "context"

func (cs *CharityService) GetCategories(ctx context.Context) ([]string, error) {
	return cs.categoriesRepository.Get(ctx)
}

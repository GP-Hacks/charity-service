package service_provider

import (
	"github.com/GP-Hacks/charity/internal/repositories/categories_repository"
	"github.com/GP-Hacks/charity/internal/repositories/collections_repository"
)

func (s *ServiceProvider) CategoriesRepository() *categories_repository.CategoriesRepository {
	if s.categoriesRepository == nil {
		s.categoriesRepository = categories_repository.NewCategoriesRepository(s.DB())
	}

	return s.categoriesRepository
}

func (s *ServiceProvider) CollectionsRepository() *collections_repository.CollectionRepository {
	if s.collectionsRepository == nil {
		s.collectionsRepository = collections_repository.NewCollectionRepository(s.DB())
	}

	return s.collectionsRepository
}

package service_provider

import "github.com/GP-Hacks/charity/internal/services/charity_service"

func (s *ServiceProvider) CharityService() *charity_service.CharityService {
	if s.chatiryService == nil {
		s.chatiryService = charity_service.NewCharityService(s.CategoriesRepository(), s.CollectionsRepository(), s.UserAdapter())
	}

	return s.chatiryService
}

package service_provider

import "github.com/GP-Hacks/charity/internal/adapters/user_adapter"

func (s *ServiceProvider) UserAdapter() *user_adapter.UserAdapter {
	if s.userAdapter == nil {
		s.userAdapter = user_adapter.NewUserAdapter()
	}

	return s.userAdapter
}

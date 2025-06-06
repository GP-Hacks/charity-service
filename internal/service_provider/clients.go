package service_provider

import (
	"github.com/GP-Hacks/proto/pkg/api/user"
)

func (s *ServiceProvider) UserClient() user.UserServiceClient {
	if s.usersClient == nil {
		s.usersClient = user.NewUserServiceClient(s.UserConnection())
	}

	return s.usersClient
}

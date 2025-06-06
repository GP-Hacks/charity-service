package user_adapter

import (
	desc "github.com/GP-Hacks/proto/pkg/api/user"
)

type UserAdapter struct {
	client desc.UserServiceClient
}

func NewUserAdapter(c desc.UserServiceClient) *UserAdapter {
	return &UserAdapter{
		client: c,
	}
}

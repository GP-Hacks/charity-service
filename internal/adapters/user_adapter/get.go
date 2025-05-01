package user_adapter

import (
	"context"

	"github.com/GP-Hacks/charity/internal/models"
)

func (ua *UserAdapter) GetByToken(ctx context.Context, token string) (*models.User, error) {
	return &models.User{
		Status: models.Admin,
	}, nil
}

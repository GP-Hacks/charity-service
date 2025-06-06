package user_adapter

import (
	"context"

	"github.com/GP-Hacks/charity/internal/models"
	"github.com/GP-Hacks/charity/internal/services"
	"github.com/GP-Hacks/proto/pkg/api/user"
	"github.com/rs/zerolog/log"
)

func (ua *UserAdapter) GetByToken(ctx context.Context, token string) (*models.User, error) {
	resp, err := ua.client.GetMe(ctx, &user.GetMeRequest{
		Token: token,
	})
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	var st models.UserStatus
	if resp.Status == user.UserStatus_ADMIN {
		st = models.Admin
	} else if resp.Status == user.UserStatus_DEFAULT {
		st = models.Default
	} else {
		return nil, services.InternalServerError
	}

	return &models.User{
		ID:          resp.Id,
		Email:       resp.User.Email,
		FirstName:   resp.User.FirstName,
		LastName:    resp.User.LastName,
		Surname:     resp.User.Surname,
		DateOfBirth: resp.User.DateOfBirth.AsTime(),
		Status:      st,
	}, nil
}

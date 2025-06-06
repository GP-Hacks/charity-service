package service_provider

import (
	"github.com/GP-Hacks/charity/internal/adapters/user_adapter"
	contrl "github.com/GP-Hacks/charity/internal/controllers/grpc"
	"github.com/GP-Hacks/charity/internal/repositories/categories_repository"
	"github.com/GP-Hacks/charity/internal/repositories/collections_repository"
	"github.com/GP-Hacks/charity/internal/services/charity_service"
	"github.com/GP-Hacks/proto/pkg/api/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

type ServiceProvider struct {
	db              *pgxpool.Pool
	usersConnection *grpc.ClientConn
	usersClient     user.UserServiceClient

	userAdapter *user_adapter.UserAdapter

	categoriesRepository  *categories_repository.CategoriesRepository
	collectionsRepository *collections_repository.CollectionRepository

	chatiryService *charity_service.CharityService

	charityController *contrl.CharityController
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

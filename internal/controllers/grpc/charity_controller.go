package grpc

import (
	"context"
	"errors"

	"github.com/GP-Hacks/charity/internal/models"
	"github.com/GP-Hacks/charity/internal/services"
	"github.com/GP-Hacks/charity/internal/services/charity_service"
	desc "github.com/GP-Hacks/proto/pkg/api/charity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CharityController struct {
	desc.UnimplementedCharityServiceServer
	charityService *charity_service.CharityService
}

func NewCharityController(cs *charity_service.CharityService) *CharityController {
	return &CharityController{
		charityService: cs,
	}
}

func (c *CharityController) GetCollections(ctx context.Context, req *desc.GetCollectionsRequest) (*desc.GetCollectionsResponse, error) {
	select {
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "Request was cancelled")
	default:
	}

	var (
		colls []*models.CollectionWithCategory
		err   error
	)

	if req.GetCategory() == "all" {
		colls, err = c.charityService.GetCollections(ctx, req.GetOffset(), req.GetLimit())
	} else {
		colls, err = c.charityService.GetCollectionsByCategory(ctx, req.GetCategory(), req.GetOffset(), req.GetLimit())
	}

	if err != nil {
		if errors.Is(err, services.NotFound) {
			return nil, status.Error(codes.NotFound, "Category not found")
		} else {
			return nil, status.Error(codes.Internal, "Internal server error")
		}
	}

	res := make([]*desc.Collection, len(colls))
	for i, coll := range colls {
		res[i] = &desc.Collection{
			Id:           int32(coll.ID),
			Category:     coll.Category,
			Name:         coll.Name,
			Description:  coll.Description,
			Organization: coll.Organization,
			Phone:        coll.Phone,
			Website:      coll.Website,
			Goal:         int32(coll.Goal),
			Current:      int32(coll.Current),
			Photo:        coll.Photo,
		}
	}

	return &desc.GetCollectionsResponse{
		Response: res,
	}, nil
}

func (c *CharityController) GetCategories(ctx context.Context, req *desc.GetCategoriesRequest) (*desc.GetCategoriesResponse, error) {
	select {
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "Request was cancelled")
	default:
	}

	cts, err := c.charityService.GetCategories(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &desc.GetCategoriesResponse{
		Categories: cts,
	}, nil
}

func (c *CharityController) Donate(ctx context.Context, req *desc.DonateRequest) (*desc.DonateResponse, error) {
	return nil, status.Error(codes.Unavailable, "Unavailable")
}

func (c *CharityController) HealthCheck(ctx context.Context, req *desc.HealthCheckRequest) (*desc.HealthCheckResponse, error) {
	return &desc.HealthCheckResponse{IsHealthy: true}, nil
}

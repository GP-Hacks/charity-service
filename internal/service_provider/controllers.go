package service_provider

import "github.com/GP-Hacks/charity/internal/controllers/grpc"

func (s *ServiceProvider) CharityController() *grpc.CharityController {
	if s.charityController == nil {
		s.charityController = grpc.NewCharityController(s.CharityService())
	}

	return s.charityController
}

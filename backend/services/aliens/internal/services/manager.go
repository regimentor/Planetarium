package services

import (
	"context"
	"github.com/regimentor/planetarium/backend/services/aliens/pkg/api"
)

type ManagerService struct {
	api.UnimplementedManagerServiceServer
}

func NewManagerService() *ManagerService {
	return &ManagerService{}
}

func (m ManagerService) Create(ctx context.Context, request *api.CreateAlienRequest) (*api.CreateAlienResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m ManagerService) Update(ctx context.Context, request *api.UpdateAlienRequest) (*api.UpdateAlienResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m ManagerService) Find(ctx context.Context, request *api.FindAlienRequest) (*api.FindAlienResponse, error) {
	//TODO implement me
	panic("implement me")
}

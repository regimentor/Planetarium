package services

import (
	"context"

	"github.com/regimentor/planetarium/backend/services/aliens/internal"
	"github.com/regimentor/planetarium/backend/services/aliens/pkg/api"
)

type AliensRepository interface {
	Create(ctx context.Context, dto *internal.CreateAlienDto) (*internal.Alien, error)
	Update(ctx context.Context, dto *internal.UpdateAlienDto) (*internal.Alien, error)
	GetByUsername(ctx context.Context, username string) (*internal.Alien, error)
}

type ManagerService struct {
	api.UnimplementedManagerServiceServer

	repository AliensRepository
}

func NewManagerService(repository AliensRepository) *ManagerService {
	return &ManagerService{
		repository: repository,
	}
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

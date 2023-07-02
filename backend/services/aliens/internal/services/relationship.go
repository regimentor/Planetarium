package services

import (
	"context"

	"github.com/regimentor/planetarium/backend/services/aliens/internal"

	"github.com/regimentor/planetarium/backend/services/aliens/pkg/api"
)

type RelationshipsRepository interface {
	Create(ctx context.Context, dto *internal.CreateRelationshipDto) (*internal.Relationship, error)
	Update(ctx context.Context, dto *internal.UpdateRelationshipDto) (*internal.Relationship, error)
}

type RelationService struct {
	api.UnimplementedRelationServiceServer

	repository RelationshipsRepository
}

func NewRelationService(repository RelationshipsRepository) *RelationService {
	return &RelationService{
		repository: repository,
	}
}

func (r RelationService) Create(ctx context.Context, request *api.CreateRelationRequestRequest) (*api.CreateRelationRequestResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RelationService) Submit(ctx context.Context, request *api.SubmitRelationRequest) (*api.SubmitRelationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RelationService) Decline(ctx context.Context, request *api.DeclineRelationRequest) (*api.DeclineRelationResponse, error) {
	//TODO implement me
	panic("implement me")
}

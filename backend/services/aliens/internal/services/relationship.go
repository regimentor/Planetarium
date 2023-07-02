package services

import (
	"context"

	"github.com/regimentor/planetarium/backend/services/aliens/pkg/api"
)

type RelationService struct {
	api.UnimplementedRelationServiceServer
}

func NewRelationService() *RelationService {
	return &RelationService{}
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

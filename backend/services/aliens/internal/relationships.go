package internal

import "context"

type RelationshipID int64

type Relationship struct {
	ID       RelationshipID `json:"id"`
	AlienID  AlienID        `json:"user_id"`
	TargetID AlienID        `json:"target_id"`
	Status   int            `json:"status"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateRelationshipDto struct {
	AlienID  AlienID `json:"user_id"`
	TargetID AlienID `json:"target_id"`
	Status   int     `json:"status"`
}

type UpdateRelationshipDto struct {
	Status int `json:"status"`
}

type RelationshipsStorage interface {
	Create(ctx context.Context, dto *CreateRelationshipDto) (*Relationship, error)
	Update(ctx context.Context, id RelationshipID, dto *UpdateRelationshipDto) (*Relationship, error)
}

type RelationshipsRepository struct {
	storage RelationshipsStorage
}

func NewRelationshipsRepository(storage RelationshipsStorage) *RelationshipsRepository {
	return &RelationshipsRepository{storage: storage}
}

func (r *RelationshipsRepository) Create(ctx context.Context, dto *CreateRelationshipDto) (*Relationship, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RelationshipsRepository) Update(ctx context.Context, dto *UpdateRelationshipDto) (*Relationship, error) {
	//TODO implement me
	panic("implement me")
}

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/regimentor/planetarium/backend/services/aliens/internal"
)

type RelationshipsStorage struct {
	con *pgxpool.Pool
}

func NewRelationshipsStorage(con *pgxpool.Pool) *RelationshipsStorage {
	return &RelationshipsStorage{con: con}
}

func (r RelationshipsStorage) Create(ctx context.Context, dto *internal.CreateRelationshipDto) (*internal.Relationship, error) {
	//TODO implement me
	panic("implement me")
}

func (r RelationshipsStorage) Update(ctx context.Context, id internal.RelationshipID, dto *internal.UpdateRelationshipDto) (*internal.Relationship, error) {
	//TODO implement me
	panic("implement me")
}

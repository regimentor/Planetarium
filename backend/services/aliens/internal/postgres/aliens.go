package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/regimentor/planetarium/backend/services/aliens/internal"
)

type AliensStorage struct {
	con *pgxpool.Pool
}

func NewAliensStorage(con *pgxpool.Pool) *AliensStorage {
	return &AliensStorage{con: con}
}

func (a AliensStorage) Create(ctx context.Context, alien *internal.CreateAlienDto) (*internal.Alien, error) {
	//TODO implement me
	panic("implement me")
}

func (a AliensStorage) Update(ctx context.Context, id internal.AlienID, dto *internal.UpdateAlienDto) (*internal.Alien, error) {
	//TODO implement me
	panic("implement me")
}

func (a AliensStorage) GetByID(ctx context.Context, id internal.AlienID) (*internal.Alien, error) {
	//TODO implement me
	panic("implement me")
}

func (a AliensStorage) GetByUsername(ctx context.Context, username string) (*internal.Alien, error) {
	//TODO implement me
	panic("implement me")
}

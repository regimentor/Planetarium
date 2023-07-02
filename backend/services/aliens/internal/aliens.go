package internal

import "context"

type AlienID int64

type Alien struct {
	ID          AlienID `json:"id"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	Password    string  `json:"-"`
	Description string  `json:"description"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateAlienDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateAlienDto struct {
	Description string `json:"description"`
}

type AliensStorage interface {
	Create(ctx context.Context, alien *CreateAlienDto) (*Alien, error)
	Update(ctx context.Context, id AlienID, dto *UpdateAlienDto) (*Alien, error)
	GetByID(ctx context.Context, id AlienID) (*Alien, error)
	GetByUsername(ctx context.Context, username string) (*Alien, error)
}

type AliensRepository struct {
	storage AliensStorage
}

func NewAliensRepository(storage AliensStorage) *AliensRepository {
	return &AliensRepository{storage: storage}
}

func (a *AliensRepository) Create(ctx context.Context, dto *CreateAlienDto) (*Alien, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AliensRepository) Update(ctx context.Context, dto *UpdateAlienDto) (*Alien, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AliensRepository) GetByUsername(ctx context.Context, username string) (*Alien, error) {
	//TODO implement me
	panic("implement me")
}

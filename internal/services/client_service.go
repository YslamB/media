package services

import (
	"context"
	"media/internal/models"
	"media/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientService struct {
	repo *repositories.ClientRepository
}

func NewClientService(db *pgxpool.Pool) *ClientService {
	return &ClientService{repositories.NewClientRepository(db)}

}

func (us *ClientService) GetUsers(ctx context.Context, id int) (int, error) {
	return 1, nil

}
func (us *ClientService) Films(ctx context.Context, page, limit int) models.Response {
	return us.repo.Films(ctx, page, limit)
}

func (us *ClientService) Books(ctx context.Context, page, limit int) models.Response {
	return us.repo.Books(ctx, page, limit)
}

func (us *ClientService) Musics(ctx context.Context, page, limit int) models.Response {
	return us.repo.Musics(ctx, page, limit)
}

func (us *ClientService) Categories(ctx context.Context) models.Response {
	return us.repo.Categories(ctx)
}

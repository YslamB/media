package repositories

import (
	"context"
	"media/internal/models"
	"media/internal/queries"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientRepository struct {
	db *pgxpool.Pool
}

func NewClientRepository(db *pgxpool.Pool) *ClientRepository {
	return &ClientRepository{db}
}

func (r *ClientRepository) GetUsers(ctx context.Context, page, limit int) (int, error) {
	// offset := page*limit - limit

	// var users []response.User
	// err := pgxscan.Select(
	// 	ctx, r.db,
	// 	&users, queries.GetUsers, offset, limit,
	// )

	return 1, nil
}

func (r *ClientRepository) Films(ctx context.Context, page, limit int) ([]models.Film, error) {
	offset := page*limit - limit
	var data []models.Film
	err := pgxscan.Select(ctx, r.db, &data, queries.GetFilms, offset, limit)

	return data, err
}

func (r *ClientRepository) Books(ctx context.Context, page, limit int) ([]models.Book, error) {
	offset := page*limit - limit
	var data []models.Book
	err := pgxscan.Select(ctx, r.db, &data, queries.GetBooks, offset, limit)

	return data, err
}

func (r *ClientRepository) Musics(ctx context.Context, page, limit int) ([]models.Music, error) {
	offset := page*limit - limit
	var data []models.Music
	err := pgxscan.Select(ctx, r.db, &data, queries.GetMusics, offset, limit)

	return data, err
}

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

func (r *ClientRepository) Films(ctx context.Context, page, limit int) models.Response {
	offset := page*limit - limit
	var data = make([]models.Film, 0)
	err := pgxscan.Select(ctx, r.db, &data, queries.GetFilms, offset, limit)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: data}
}

func (r *ClientRepository) Books(ctx context.Context, page, limit int) models.Response {
	offset := page*limit - limit
	var data = make([]models.Book, 0)
	err := pgxscan.Select(ctx, r.db, &data, queries.GetBooks, offset, limit)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: data}
}

func (r *ClientRepository) Musics(ctx context.Context, page, limit int) models.Response {
	offset := page*limit - limit
	var data []models.Music
	err := pgxscan.Select(ctx, r.db, &data, queries.GetMusics, offset, limit)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: data}
}

func (r *ClientRepository) Categories(ctx context.Context) models.Response {

	var data []models.Categories
	err := pgxscan.Select(ctx, r.db, &data, queries.GetCategories)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: data}

}

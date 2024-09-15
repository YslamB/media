package repositories

import (
	"context"

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

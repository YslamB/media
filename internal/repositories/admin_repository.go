package repositories

import (
	"context"
	"fmt"
	"media/internal/models"
	"media/internal/queries"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminRepository struct {
	db *pgxpool.Pool
}

func NewAdminRepository(db *pgxpool.Pool) *AdminRepository {
	return &AdminRepository{db}
}

func (r *AdminRepository) File(ctx context.Context, path, title, desc, fileType string) (int, error) {
	return 1, nil
}

func (r *AdminRepository) Music(ctx context.Context, path, title, desc, language string) (int, error) {
	return 1, nil
}

func (r *AdminRepository) Film(ctx context.Context, path, title, desc, language string) (int, error) {
	return 1, nil
}

func (r *AdminRepository) GetAdmin(ctx context.Context, username string) models.Admin {
	var admin models.Admin

	err := r.db.QueryRow(
		ctx, queries.GetAdmin, username,
	).Scan(&admin.Username, &admin.Password)
	fmt.Println(err)

	return admin
}

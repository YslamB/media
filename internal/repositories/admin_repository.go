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

func (r *AdminRepository) GetMusicPath(ctx context.Context, id string) string {
	var path string
	r.db.QueryRow(ctx, queries.DeleteMusic, id).Scan(&path)
	return path
}

func (r *AdminRepository) GetFilmPath(ctx context.Context, id string) string {
	var path string
	r.db.QueryRow(ctx, queries.DeleteFilm, id).Scan(&path)
	return path
}

func (r *AdminRepository) GetBookPath(ctx context.Context, id string) string {
	var path string
	r.db.QueryRow(ctx, queries.DeleteBook, id).Scan(&path)
	return path
}

func (r *AdminRepository) Music(ctx context.Context, path, imagePath, title, desc, language, categoryId string) (string, error) {
	var id string
	err := r.db.QueryRow(ctx, queries.CreateMusic, categoryId, language, title, desc, path, imagePath).Scan(&id)
	return id, err
}

func (r *AdminRepository) Film(ctx context.Context, title, path, imagePath, desc, language, categoryId string) (string, error) {
	var id string
	err := r.db.QueryRow(ctx, queries.CreateFilm, categoryId, language, title, desc, path, imagePath).Scan(&id)
	return id, err
}

func (r *AdminRepository) Book(ctx context.Context, path, imagePath, title, desc, language, categoryId string) (string, error) {
	var id string
	err := r.db.QueryRow(ctx, queries.CreateBook, categoryId, language, title, desc, path, imagePath).Scan(&id)
	return id, err
}

func (r *AdminRepository) GetAdmin(ctx context.Context, username string) models.Admin {
	var admin models.Admin

	err := r.db.QueryRow(
		ctx, queries.GetAdmin, username,
	).Scan(&admin.Username, &admin.Password)
	fmt.Println(err)

	return admin
}

func (r *AdminRepository) CreateCategory(ctx context.Context, ctg models.Category) (int, error) {

	jsonData := fmt.Sprintf(`{"ru": "%s", "tm": "%s"}`, ctg.Ru, ctg.Tm)
	var id int
	err := r.db.QueryRow(ctx, queries.CreateCategory, jsonData).Scan(&id)
	return id, err

}

func (r *AdminRepository) CreateSubCategory(ctx context.Context, ctg models.Category) (int, error) {

	jsonData := fmt.Sprintf(`{"ru": "%s", "tm": "%s"}`, ctg.Ru, ctg.Tm)
	var id int
	fmt.Println(ctg.ID)
	err := r.db.QueryRow(ctx, queries.CreateSubCategory, ctg.ID, jsonData).Scan(&id)
	return id, err

}

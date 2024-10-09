package repositories

import (
	"context"
	"fmt"
	"media/internal/models"
	"media/internal/queries"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminRepository struct {
	DB *pgxpool.Pool
}

func NewAdminRepository(db *pgxpool.Pool) *AdminRepository {
	return &AdminRepository{db}
}

func (r *AdminRepository) File(ctx context.Context, path, title, desc, fileType string) (int, error) {
	return 1, nil
}

func (r *AdminRepository) GetMusicPath(ctx context.Context, id string) string {
	var path string
	r.DB.QueryRow(ctx, queries.DeleteMusic, id).Scan(&path)
	return path
}

func (r *AdminRepository) GetFilmPath(ctx context.Context, id string) string {
	var path string
	r.DB.QueryRow(ctx, queries.DeleteFilm, id).Scan(&path)
	return path
}

func (r *AdminRepository) DeleteBook(ctx context.Context, id string) string {
	var path string
	r.DB.QueryRow(ctx, queries.DeleteBook, id).Scan(&path)
	return path
}

func (r *AdminRepository) Music(ctx context.Context, path, imagePath, title, desc, language, categoryId string) (string, error) {
	var id string
	err := r.DB.QueryRow(ctx, queries.CreateMusic, categoryId, language, title, desc, path, imagePath).Scan(&id)
	return id, err
}

func (r *AdminRepository) Film(ctx context.Context, title, desc, language string, categoryId int) (string, error) {
	var id string
	err := r.DB.QueryRow(ctx, queries.CreateFilm, categoryId, language, title, desc).Scan(&id)
	return id, err
}

func (r *AdminRepository) Book(ctx context.Context, title, desc, language string, categoryId int) (string, error) {
	var id string
	err := r.DB.QueryRow(ctx, queries.CreateBook, categoryId, language, title, desc).Scan(&id)
	return id, err
}

func (r *AdminRepository) GetAdmin(ctx context.Context, username string) models.Admin {
	var admin models.Admin

	r.DB.QueryRow(
		ctx, queries.GetAdmin, username,
	).Scan(&admin.Username, &admin.Password)

	return admin
}

func (r *AdminRepository) CreateCategory(ctx context.Context, ctg models.Category) (int, error) {

	jsonData := fmt.Sprintf(`{"ru": "%s", "tm": "%s"}`, ctg.Ru, ctg.Tm)
	var id int
	err := r.DB.QueryRow(ctx, queries.CreateCategory, jsonData).Scan(&id)
	return id, err

}

func (r *AdminRepository) CreateSubCategory(ctx context.Context, ctg models.Category) (int, error) {

	jsonData := fmt.Sprintf(`{"ru": "%s", "tm": "%s"}`, ctg.Ru, ctg.Tm)
	var id int
	err := r.DB.QueryRow(ctx, queries.CreateSubCategory, ctg.ID, jsonData).Scan(&id)
	return id, err

}

func (r *AdminRepository) UpdateBook(ctx context.Context, title, description, language string, categoryId, bookID int) error {
	var filePath, imagePath = "", ""
	err := r.DB.QueryRow(ctx, queries.UpdateBook, categoryId, language, title, description, bookID).Scan(&filePath, &imagePath)

	return err
}

func (r *AdminRepository) UpdateFilm(ctx context.Context, title, description, language string, filmID, categoryId int) error {
	var filePath, imagePath = "", ""
	err := r.DB.QueryRow(ctx, queries.UpdateFilm, categoryId, language, title, description, filmID).Scan(&filePath, &imagePath)

	return err
}

func (r *AdminRepository) UpdateMusic(ctx context.Context, title, description, language string, categoryId, musicID int) error {
	var filePath, imagePath = "", ""
	err := r.DB.QueryRow(ctx, queries.UpdateMusic, categoryId, language, title, description, musicID).Scan(&filePath, &imagePath)

	return err
}

func (r *AdminRepository) GetFilmImageFilePath(ctx context.Context, id int) (string, string, int) {
	var filePath, imagePath = "", ""
	fmt.Println(id)
	err := r.DB.QueryRow(ctx, queries.GetImageFilePathFilm, id).Scan(&filePath, &imagePath, &id)

	if err != nil {
		fmt.Println(err)
		return "", "", 0
	}

	return filePath, imagePath, id
}

func (r *AdminRepository) GetBookImageFilePath(ctx context.Context, id int) (string, string, int) {
	var filePath, imagePath = "", ""

	err := r.DB.QueryRow(ctx, queries.GetImageFilePathBook, id).Scan(&filePath, &imagePath, &id)

	if err != nil {
		fmt.Println(err)
		return "", "", 0
	}

	return filePath, imagePath, id
}

func (r *AdminRepository) GetMusicImageFilePath(ctx context.Context, id int) (string, string, int) {
	var filePath, imagePath = "", ""

	err := r.DB.QueryRow(ctx, queries.GetImageFilePathMusic, id).Scan(&filePath, &imagePath, &id)

	if err != nil {
		fmt.Println(err)
		return "", "", 0
	}

	return filePath, imagePath, id
}

func (r *AdminRepository) UpdateFilmImage(ctx context.Context, path string, id int) {
	fmt.Println(id)
	r.DB.Exec(ctx, queries.UpdateFilmImage, path, id)
}
func (r *AdminRepository) UpdateBookImage(ctx context.Context, path string, id int) {
	fmt.Println(id)
	r.DB.Exec(ctx, queries.UpdateFilmImage, path, id)
}

func (r *AdminRepository) UpdateFilmPath(ctx context.Context, path string, id int) {
	fmt.Println(id)
	r.DB.Exec(ctx, queries.UpdateFilmPath, path, id)
}

func (r *AdminRepository) UpdateBookPath(ctx context.Context, path string, id int) {
	_, err := r.DB.Exec(ctx, queries.UpdateBookPath, path, id)
	fmt.Println("update book path error: ", path)
	fmt.Println(err)
}

func (r *AdminRepository) UpdateMusicImage(ctx context.Context, path string, id int) {
	fmt.Println(id)
	r.DB.Exec(ctx, queries.UpdateMusicImage, path, id)
}

func (r *AdminRepository) UpdateMusicPath(ctx context.Context, path string, id int) {
	fmt.Println(id)
	r.DB.Exec(ctx, queries.UpdateMusicPath, path, id)
}

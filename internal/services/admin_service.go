package services

import (
	"context"
	"errors"
	"fmt"
	"media/internal/models"
	"media/internal/repositories"
	"media/pkg/utils"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	repo *repositories.AdminRepository
}

func NewAdminService(db *pgxpool.Pool) *AdminService {
	return &AdminService{repositories.NewAdminRepository(db)}
}

func (us *AdminService) GetUsers(ctx context.Context, id int) models.Response {

	return models.Response{Data: "users"}

}

func (sr *AdminService) DeleteMusic(ctx context.Context, id string) models.Response {
	path := sr.repo.GetMusicPath(ctx, id)
	if path == "" {
		return models.Response{Error: errors.New("not found"), Status: 404}
	}
	os.RemoveAll(filepath.Dir(path))
	return models.Response{Data: gin.H{"message": "deleted"}}
}

func (sr *AdminService) Music(ctx context.Context, form *multipart.Form) models.Response {

	musics := form.File["music"]
	images := form.File["image"]

	if len(musics) == 0 || len(images) == 0 {
		return models.Response{Error: errors.New("no musics or images found in the request"), Status: 400}
	}

	musicExt := filepath.Ext(musics[0].Filename)
	imageEXT := filepath.Ext(images[0].Filename)

	if musicExt != ".mp3" || imageEXT != ".jpg" {
		return models.Response{Error: errors.New("invalid file type, must be  .mp3 and .jpg "), Status: 400}
	}

	timestamp := time.Now().Unix()
	musicFilename := fmt.Sprintf("%d%s", timestamp, musicExt)
	imageFilename := fmt.Sprintf("%d%s", timestamp, imageEXT)
	title := form.Value["title"]
	categoryId := form.Value["category_id"]
	description := form.Value["description"]
	language := form.Value["language"]
	uploadMusicFilePath := fmt.Sprintf("./uploads/music/%d/", timestamp)
	err := utils.SaveUploadedFile(musics[0], uploadMusicFilePath+musicFilename)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	err = utils.SaveUploadedFile(images[0], uploadMusicFilePath+imageFilename)

	if err != nil {
		os.RemoveAll(uploadMusicFilePath)
		return models.Response{Error: err, Status: 500}
	}

	err = utils.ResizeImage(uploadMusicFilePath+imageFilename, 700)

	if err != nil {
		os.RemoveAll(uploadMusicFilePath)
		return models.Response{Error: err, Status: 500}
	}

	id, err := sr.repo.Music(ctx, uploadMusicFilePath+fmt.Sprint(timestamp)+"HLS.m3u8",
		uploadMusicFilePath+imageFilename, title[0], description[0], language[0], categoryId[0])

	if err == nil {
		go utils.ConvertToHLS(uploadMusicFilePath, musicFilename, "music")
	} else {
		os.RemoveAll(uploadMusicFilePath)
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: &gin.H{"id": id}}
}

func (sr *AdminService) DeleteFilm(ctx context.Context, id string) models.Response {
	path := sr.repo.GetFilmPath(ctx, id)

	if path == "" {
		return models.Response{Error: errors.New("not found"), Status: 404}
	}

	err := os.RemoveAll(filepath.Dir(path))

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: gin.H{"message": "deleted"}}
}

func (sr *AdminService) Film(ctx context.Context, form *multipart.Form) models.Response {

	films := form.File["film"]
	images := form.File["image"]

	if len(films) == 0 || len(images) == 0 {
		return models.Response{Error: errors.New("no films or images found in the request"), Status: 400}
	}

	filmExt := filepath.Ext(films[0].Filename)
	imageEXT := filepath.Ext(images[0].Filename)

	if filmExt != ".mp4" || imageEXT != ".jpg" {
		return models.Response{Error: errors.New("invalid file type, must be  .mp3 and .jpg "), Status: 400}
	}

	timestamp := time.Now().Unix()
	filmFilename := fmt.Sprintf("%d%s", timestamp, filmExt)
	imageFilename := fmt.Sprintf("%d%s", timestamp, imageEXT)
	title := form.Value["title"]
	description := form.Value["description"]
	categoryId := form.Value["category_id"]
	language := form.Value["language"]
	uploadfilmFilePath := fmt.Sprintf("./uploads/film/%d/", timestamp)
	err := utils.SaveUploadedFile(films[0], uploadfilmFilePath+filmFilename)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	err = utils.SaveUploadedFile(images[0], uploadfilmFilePath+imageFilename)

	if err != nil {
		os.RemoveAll(uploadfilmFilePath)
		return models.Response{Error: err, Status: 500}
	}

	err = utils.ResizeImage(uploadfilmFilePath+imageFilename, 700)

	if err != nil {
		os.RemoveAll(uploadfilmFilePath)
		return models.Response{Error: err, Status: 500}
	}

	id, err := sr.repo.Film(ctx, title[0], uploadfilmFilePath+fmt.Sprint(timestamp)+"HLS.m3u8",
		uploadfilmFilePath+imageFilename, description[0], language[0], categoryId[0])

	if err == nil {
		go utils.ConvertToHLS(uploadfilmFilePath, filmFilename, "film")
	} else {
		os.RemoveAll(uploadfilmFilePath)
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: &gin.H{"id": id}}
}

func (sr *AdminService) DeleteBook(ctx context.Context, id string) models.Response {
	path := sr.repo.GetBookPath(ctx, id)

	if path == "" {
		return models.Response{Error: errors.New("not found"), Status: 404}
	}

	os.RemoveAll(filepath.Dir(path))

	return models.Response{Data: gin.H{"message": "deleted"}}
}

func (sr *AdminService) Book(ctx context.Context, form *multipart.Form) models.Response {

	books := form.File["book"]
	images := form.File["image"]

	if len(books) == 0 || len(images) == 0 {
		return models.Response{Error: errors.New("no books or images found in the request"), Status: 400}
	}

	bookExt := filepath.Ext(books[0].Filename)
	imageEXT := filepath.Ext(images[0].Filename)

	if bookExt != ".pdf" || imageEXT != ".jpg" {
		return models.Response{Error: errors.New("invalid file type, must be  .pdf and .jpg "), Status: 400}
	}

	timestamp := time.Now().Unix()
	bookFilename := fmt.Sprintf("%d%s", timestamp, bookExt)
	imageFilename := fmt.Sprintf("%d%s", timestamp, imageEXT)
	title := form.Value["title"]
	description := form.Value["description"]
	language := form.Value["language"]
	categoryId := form.Value["category_id"]
	uploadbookFilePath := fmt.Sprintf("./uploads/book/%d/", timestamp)
	err := utils.SaveUploadedFile(books[0], uploadbookFilePath+bookFilename)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	err = utils.SaveUploadedFile(images[0], uploadbookFilePath+imageFilename)

	if err != nil {
		os.RemoveAll(uploadbookFilePath)
		return models.Response{Error: err, Status: 500}
	}

	err = utils.ResizeImage(uploadbookFilePath+imageFilename, 700)

	if err != nil {
		os.RemoveAll(uploadbookFilePath)
		return models.Response{Error: err, Status: 500}
	}

	id, err := sr.repo.Book(ctx, uploadbookFilePath+bookFilename, uploadbookFilePath+imageFilename,
		title[0], description[0], language[0], categoryId[0])

	if err != nil {
		os.RemoveAll(uploadbookFilePath)
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: &gin.H{"id": id}}
}

func (sr *AdminService) AdminLogin(ctx context.Context, admin models.LoginForm) (string, string, error) {

	findedAdmin := sr.repo.GetAdmin(ctx, admin.Username)
	compareError := bcrypt.CompareHashAndPassword(
		[]byte(findedAdmin.Password), []byte(admin.Password),
	)

	if compareError != nil {
		return "", "", compareError
	}

	accessToken, refreshToken := utils.CreateRefreshAccsessToken(findedAdmin.Username, "admin")
	return accessToken, refreshToken, nil
}

func (sr *AdminService) Category(ctx context.Context, ctg models.Category) models.Response {

	id, err := sr.repo.CreateCategory(ctx, ctg)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: gin.H{"id": id}}

}

func (sr *AdminService) SubCategory(ctx context.Context, ctg models.Category) models.Response {

	id, err := sr.repo.CreateSubCategory(ctx, ctg)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: gin.H{"id": id}}

}

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

func (us *AdminService) GetUsers(ctx context.Context, id int) (int, error) {

	// otp := response.OTP{}
	// userPhone := us.repo.CheckUserExist(ctx, user.Phone)

	// if userPhone != "" {
	// 	return otp
	// }

	return 1, nil

}

func (sr *AdminService) File(ctx context.Context, form *multipart.Form) (any, error) {

	files := form.File["file"]

	if len(files) == 0 {
		return nil, errors.New("no files found in the request")
	}

	ext := filepath.Ext(files[0].Filename)
	validExtensions := map[string]bool{".mp4": true, ".mp3": true, ".pdf": true}

	if !validExtensions[ext] {
		return nil, errors.New("invalid file type, must be .mp4, .mp3 or .pdf")
	}

	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%d%s", timestamp, ext)
	title := form.Value["title"]
	description := form.Value["description"]
	contentType := files[0].Header.Get("Content-Type")
	fileType := utils.GetType(contentType)
	uploadFilePath := fmt.Sprintf("./uploads/%s/%d/", fileType, timestamp)
	err := utils.SaveUploadedFile(files[0], uploadFilePath+filename)

	if err != nil {
		return nil, err
	}

	switch fileType {
	case "video":
		go utils.ConvertToHLS(uploadFilePath, filename, "video")
	case "audio":
		go utils.ConvertToHLS(uploadFilePath, filename, "audio")
	}

	id, err := sr.repo.File(ctx, uploadFilePath+filename, title[0], description[0], fileType)

	return &gin.H{"id": id}, err
}

func (sr *AdminService) DeleteMusic(ctx context.Context, id string) error {
	path := sr.repo.GetMusicPath(ctx, id)
	if path == "" {
		return errors.New("not found")
	}
	os.RemoveAll(filepath.Dir(path))
	return nil
}

func (sr *AdminService) Music(ctx context.Context, form *multipart.Form) (any, error) {

	musics := form.File["music"]
	images := form.File["image"]

	if len(musics) == 0 || len(images) == 0 {
		return nil, errors.New("no musics or images found in the request")
	}

	musicExt := filepath.Ext(musics[0].Filename)
	imageEXT := filepath.Ext(images[0].Filename)

	if musicExt != ".mp3" || imageEXT != ".jpg" {
		return nil, errors.New("invalid file type, must be  .mp3 and .jpg ")
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
		return nil, err
	}

	err = utils.SaveUploadedFile(images[0], uploadMusicFilePath+imageFilename)

	if err != nil {
		os.RemoveAll(uploadMusicFilePath)
		return nil, err
	}

	err = utils.ResizeImage(uploadMusicFilePath+imageFilename, 700)

	if err != nil {
		os.RemoveAll(uploadMusicFilePath)
		return nil, err
	}

	id, err := sr.repo.Music(ctx, uploadMusicFilePath+fmt.Sprint(timestamp)+"HLS.m3u8",
		uploadMusicFilePath+imageFilename, title[0], description[0], language[0], categoryId[0])

	if err == nil {
		go utils.ConvertToHLS(uploadMusicFilePath, musicFilename, "music")
	} else {
		os.RemoveAll(uploadMusicFilePath)
	}

	return &gin.H{"id": id}, err
}

func (sr *AdminService) DeleteFilm(ctx context.Context, id string) error {
	path := sr.repo.GetFilmPath(ctx, id)
	if path == "" {
		return errors.New("not found")
	}
	os.RemoveAll(filepath.Dir(path))
	return nil
}

func (sr *AdminService) Film(ctx context.Context, form *multipart.Form) (any, error) {

	films := form.File["film"]
	images := form.File["image"]

	if len(films) == 0 || len(images) == 0 {
		return nil, errors.New("no films or images found in the request")
	}

	filmExt := filepath.Ext(films[0].Filename)
	imageEXT := filepath.Ext(images[0].Filename)

	if filmExt != ".mp4" || imageEXT != ".jpg" {
		return nil, errors.New("invalid file type, must be  .mp3 and .jpg ")
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
		return nil, err
	}

	err = utils.SaveUploadedFile(images[0], uploadfilmFilePath+imageFilename)

	if err != nil {
		os.RemoveAll(uploadfilmFilePath)
		return nil, err
	}

	err = utils.ResizeImage(uploadfilmFilePath+imageFilename, 700)

	if err != nil {
		os.RemoveAll(uploadfilmFilePath)
		return nil, err
	}

	id, err := sr.repo.Film(ctx, uploadfilmFilePath+fmt.Sprint(timestamp)+"HLS.m3u8", title[0],
		uploadfilmFilePath+imageFilename, description[0], language[0], categoryId[0])

	if err == nil {
		go utils.ConvertToHLS(uploadfilmFilePath, filmFilename, "film")
	} else {
		os.RemoveAll(uploadfilmFilePath)
	}

	return &gin.H{"id": id}, err
}

func (sr *AdminService) DeleteBook(ctx context.Context, id string) error {
	path := sr.repo.GetBookPath(ctx, id)
	if path == "" {
		return errors.New("not found")
	}
	os.RemoveAll(filepath.Dir(path))
	return nil
}

func (sr *AdminService) Book(ctx context.Context, form *multipart.Form) (any, int, error) {

	books := form.File["book"]
	images := form.File["image"]

	if len(books) == 0 || len(images) == 0 {
		return nil, 400, errors.New("no books or images found in the request")
	}

	bookExt := filepath.Ext(books[0].Filename)
	imageEXT := filepath.Ext(images[0].Filename)

	if bookExt != ".pdf" || imageEXT != ".jpg" {
		return nil, 403, errors.New("invalid file type, must be  .mp3 and .jpg ")
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
		return nil, 0, err
	}

	err = utils.SaveUploadedFile(images[0], uploadbookFilePath+imageFilename)

	if err != nil {
		os.RemoveAll(uploadbookFilePath)
		return nil, 0, err
	}

	err = utils.ResizeImage(uploadbookFilePath+imageFilename, 700)

	if err != nil {
		os.RemoveAll(uploadbookFilePath)
		return nil, 0, err
	}

	id, err := sr.repo.Book(ctx, uploadbookFilePath+bookFilename, uploadbookFilePath+imageFilename,
		title[0], description[0], language[0], categoryId[0])

	if err != nil {
		os.RemoveAll(uploadbookFilePath)
	}

	return &gin.H{"id": id}, 0, err
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

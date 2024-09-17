package services

import (
	"context"
	"errors"
	"fmt"
	"media/internal/models"
	"media/internal/repositories"
	"media/pkg/utils"
	"mime/multipart"
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
	fmt.Println(title)
	fmt.Println(description)
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

func (sr *AdminService) Music(ctx context.Context, form *multipart.Form) (any, error) {

	musics := form.File["music"]
	images := form.File["image"]

	if len(musics) == 0 || len(images) == 0 {
		return nil, errors.New("no musics or images found in the request")
	}

	musicExt := filepath.Ext(musics[0].Filename)
	imageEXT := filepath.Ext(images[0].Filename)

	fmt.Println(imageEXT)
	if musicExt != ".mp3" || imageEXT != ".jpg" {
		return nil, errors.New("invalid file type, must be  .mp3 and .jpg ")
	}

	timestamp := time.Now().Unix()
	musicFilename := fmt.Sprintf("%d%s", timestamp, musicExt)
	imageFilename := fmt.Sprintf("%d%s", timestamp, imageEXT)
	title := form.Value["title"]
	description := form.Value["description"]
	language := form.Value["language"]

	uploadMusicFilePath := fmt.Sprintf("./uploads/music/%d/", timestamp)

	err := utils.SaveUploadedFile(musics[0], uploadMusicFilePath+musicFilename)

	if err != nil {
		return nil, err
	}

	err = utils.SaveUploadedFile(images[0], uploadMusicFilePath+imageFilename)

	if err != nil {
		return nil, err
	}
	fmt.Println(uploadMusicFilePath + imageFilename)
	err = utils.ResizeImage(uploadMusicFilePath+imageFilename, 700)
	if err != nil {
		return nil, err
	}

	go utils.ConvertToHLS(uploadMusicFilePath, musicFilename, "music")

	id, err := sr.repo.Music(ctx, uploadMusicFilePath+musicFilename, title[0], description[0], language[0])

	return &gin.H{"id": id}, err
}

func (sr *AdminService) Film(ctx context.Context, form *multipart.Form) (any, error) {

	films := form.File["film"]
	images := form.File["image"]
	fmt.Println(len(films))
	fmt.Println(len(images))

	if len(films) == 0 || len(images) == 0 {
		return nil, errors.New("no films or images found in the request")
	}

	filmExt := filepath.Ext(films[0].Filename)
	imageEXT := filepath.Ext(images[0].Filename)

	fmt.Println(imageEXT)
	if filmExt != ".mp4" || imageEXT != ".jpg" {
		return nil, errors.New("invalid file type, must be  .mp3 and .jpg ")
	}

	timestamp := time.Now().Unix()
	filmFilename := fmt.Sprintf("%d%s", timestamp, filmExt)
	imageFilename := fmt.Sprintf("%d%s", timestamp, imageEXT)
	title := form.Value["title"]
	description := form.Value["description"]
	language := form.Value["language"]

	uploadfilmFilePath := fmt.Sprintf("./uploads/film/%d/", timestamp)

	err := utils.SaveUploadedFile(films[0], uploadfilmFilePath+filmFilename)

	if err != nil {
		return nil, err
	}

	err = utils.SaveUploadedFile(images[0], uploadfilmFilePath+imageFilename)

	if err != nil {
		return nil, err
	}
	fmt.Println(uploadfilmFilePath + imageFilename)
	err = utils.ResizeImage(uploadfilmFilePath+imageFilename, 700)
	if err != nil {
		return nil, err
	}

	go utils.ConvertToHLS(uploadfilmFilePath, filmFilename, "film")

	id, err := sr.repo.Film(ctx, uploadfilmFilePath+filmFilename, title[0], description[0], language[0])

	return &gin.H{"id": id}, err
}

func (sr *AdminService) AdminLogin(ctx context.Context, admin models.LoginForm) (string, string, error) {
	findedAdmin := sr.repo.GetAdmin(ctx, admin.Username)
	fmt.Println(findedAdmin.Password)
	compareError := bcrypt.CompareHashAndPassword(
		[]byte(findedAdmin.Password), []byte(admin.Password),
	)

	if compareError != nil {
		return "", "", compareError
	}

	accessToken, refreshToken := utils.CreateRefreshAccsessToken(findedAdmin.Username, "admin")
	return accessToken, refreshToken, nil
}

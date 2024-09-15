package services

import (
	"context"
	"errors"
	"fmt"
	"media/internal/repositories"
	"media/pkg/utils"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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

	if len(musics) == 0 {
		return nil, errors.New("no musics found in the request")
	}

	musicExt := filepath.Ext(musics[0].Filename)
	imageEXT := filepath.Ext(images[0].Filename)

	if imageEXT != ".mp3" {
		return nil, errors.New("invalid file type, must be .mp4, .mp3 or .pdf")
	}

	timestamp := time.Now().Unix()
	musicFilename := fmt.Sprintf("%d%s", timestamp, musicExt)
	imageFilename := fmt.Sprintf("%d%s", timestamp, imageEXT)
	title := form.Value["title"]
	description := form.Value["description"]
	language := form.Value["language"]

	contentType := musics[0].Header.Get("Content-Type")
	fileType := utils.GetType(contentType)
	uploadFilePath := fmt.Sprintf("./uploads/%s/%d/", fileType, timestamp)

	err := utils.SaveUploadedFile(musics[0], uploadFilePath+musicFilename)

	if err != nil {
		return nil, err
	}

	err = utils.SaveUploadedFile(images[0], uploadFilePath+imageFilename)

	if err != nil {
		return nil, err
	}

	go utils.ConvertToHLS(uploadFilePath, musicFilename, "audio")

	id, err := sr.repo.Music(ctx, uploadFilePath+musicFilename, title[0], description[0], fileType, language[0])

	return &gin.H{"id": id}, err
}

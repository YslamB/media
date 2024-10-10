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
	"strings"
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

func (sr *AdminService) Film(ctx context.Context, film models.ElementData) models.Response {

	id, err := sr.repo.Film(ctx, film.Title, film.Description, film.Language, film.CategoryID)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: &gin.H{"id": id}}
}

func (sr *AdminService) Music(ctx context.Context, music models.ElementData) models.Response {

	id, err := sr.repo.Music(ctx, music.Title, music.Description, music.Language, music.CategoryID)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: &gin.H{"id": id}}
}

func (sr *AdminService) DeleteBook(ctx context.Context, id string) models.Response {
	path := sr.repo.DeleteBook(ctx, id)

	if path == "" {
		return models.Response{Error: errors.New("not found"), Status: 404}
	}

	os.RemoveAll(filepath.Dir(path))

	return models.Response{Data: gin.H{"message": "deleted"}}
}

func (sr *AdminService) Book(ctx context.Context, form models.ElementData) models.Response {

	title := form.Title
	description := form.Description
	language := form.Language
	categoryId := form.CategoryID

	id, err := sr.repo.Book(ctx, title, description, language, categoryId)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: &gin.H{"id": id}}
}

// func (sr *AdminService) Book(ctx context.Context, form models.ElementData) models.Response {

// 	books := form.File["book"]
// 	images := form.File["image"]

// 	if len(books) == 0 || len(images) == 0 {
// 		return models.Response{Error: errors.New("no books or images found in the request"), Status: 400}
// 	}

// 	bookExt := filepath.Ext(books[0].Filename)
// 	imageEXT := filepath.Ext(images[0].Filename)

// 	allowedImageExts := map[string]bool{
// 		".jpg":  true,
// 		".jpeg": true,
// 		".png":  true,
// 		".webp": true,
// 	}

// 	if bookExt != ".pdf" || !allowedImageExts[imageEXT] {
// 		return models.Response{Error: errors.New("invalid file type, must be  .pdf and .jpg "), Status: 400}
// 	}

// 	timestamp := time.Now().Unix()
// 	bookFilename := fmt.Sprintf("%d%s", timestamp, bookExt)
// 	imageFilename := fmt.Sprintf("%d%s", timestamp, imageEXT)
// 	title := form.Value["title"]
// 	description := form.Value["description"]
// 	language := form.Value["language"]
// 	categoryId := form.Value["category_id"]
// 	uploadbookFilePath := fmt.Sprintf("./uploads/book/%d/", timestamp)
// 	err := utils.SaveUploadedFile(books[0], uploadbookFilePath+bookFilename)

// 	if err != nil {
// 		return models.Response{Error: err, Status: 500}
// 	}

// 	err = utils.SaveUploadedFile(images[0], uploadbookFilePath+imageFilename)

// 	if err != nil {
// 		os.RemoveAll(uploadbookFilePath)
// 		return models.Response{Error: err, Status: 500}
// 	}

// 	status, err := utils.ResizeImage(uploadbookFilePath+imageFilename, 700)

// 	if err != nil {
// 		os.RemoveAll(uploadbookFilePath)
// 		return models.Response{Error: err, Status: status}
// 	}

// 	id, err := sr.repo.Book(ctx, uploadbookFilePath[1:]+bookFilename, uploadbookFilePath[1:]+imageFilename,
// 		title[0], description[0], language[0], categoryId[0])

// 	if err != nil {
// 		os.RemoveAll(uploadbookFilePath)
// 		return models.Response{Error: err, Status: 500}
// 	}

// 	return models.Response{Data: &gin.H{"id": id}}
// }

func (sr *AdminService) UpdateFilm(ctx context.Context, form *multipart.Form, element models.ElementData, method string) models.Response {
	filmFilePath, filmImagePath, id := sr.repo.GetFilmImageFilePath(ctx, element.ID)

	if id == 0 {
		return models.Response{Error: errors.New("not found"), Status: 404}
	}

	if method == "POST" {
		id := form.Value["id"][0]

		films := form.File["film"]
		images := form.File["image"]

		if len(films) > 1 || len(images) > 1 {
			return models.Response{Error: errors.New("too many files"), Status: 400}
		}

		if len(films) == 1 {

			if filmFilePath == "" {

				if filmImagePath != "" {
					filmFilePath = strings.TrimSuffix(filmImagePath, filepath.Base(filmImagePath)) + "/hls/" + strings.TrimSuffix(filepath.Base(filmImagePath), filepath.Ext(filmImagePath)) + ".m3u8"
				} else {
					timestamp := time.Now().Unix()
					filmFilePath = fmt.Sprintf("/uploads/film/%d/hls/%d.m3u8", timestamp, timestamp)
				}
				fmt.Println("filmFilePath")
				fmt.Println(filmFilePath)
				sr.repo.UpdateFilmPath(ctx, filmFilePath, element.ID)
			}

			filmExt := filepath.Ext(films[0].Filename)

			if filmExt != ".mp4" {
				return models.Response{Error: errors.New("invalid file type, must be  .mp4 "), Status: 400}
			}

			os.RemoveAll("." + filepath.Dir(filmFilePath))
			time.Sleep(5 * time.Second)
			err := utils.SaveUploadedFile(films[0], "."+strings.TrimSuffix(filmFilePath, filepath.Ext(filmFilePath))+filmExt)

			if err != nil {
				return models.Response{Error: err, Status: 500}
			}

			go utils.ConvertToHLS("."+strings.TrimSuffix(filmFilePath, filepath.Base(filmFilePath)), strings.TrimSuffix(filepath.Base(filmFilePath), filepath.Ext(filepath.Base(filmFilePath)))+filmExt, "film")

			return models.Response{Data: &gin.H{"id": id}}

		}

		if len(images) == 1 {

			if filmImagePath == "" {

				if filmFilePath != "" {
					startIndex := strings.Index(filmFilePath, "hls")
					filmImagePath = filmFilePath[:startIndex] + strings.TrimSuffix(filepath.Base(filmFilePath), filepath.Ext(filmFilePath)) + ".webp"
				} else {
					timestamp := time.Now().Unix()
					filmImagePath = fmt.Sprintf("/uploads/film/%d/%d.webp", timestamp, timestamp)

				}
				sr.repo.UpdateFilmImage(ctx, filmImagePath, element.ID)

				fmt.Println(filmFilePath)
				fmt.Println(filmFilePath)
			}

			allowedImageExts := map[string]bool{
				".jpg":  true,
				".jpeg": true,
				".png":  true,
				".webp": true,
			}

			imageEXT := filepath.Ext(images[0].Filename)
			if !allowedImageExts[imageEXT] {
				// todo: remove if uploaded film
				return models.Response{Error: errors.New("invalid file type, must be  .pdf "), Status: 400}
			}

			os.RemoveAll("." + filmImagePath)
			err := utils.SaveUploadedFile(images[0], "."+strings.TrimSuffix(filmImagePath, filepath.Ext(filmImagePath))+imageEXT)

			if err != nil {
				return models.Response{Error: err, Status: 500}
			}

			fmt.Println("Ssss")

			status, err := utils.ResizeImage("."+strings.TrimSuffix(filmImagePath, filepath.Ext(filmImagePath))+imageEXT, 700)

			if err != nil {
				os.RemoveAll("." + filmImagePath)
				return models.Response{Error: err, Status: status}
			}

			return models.Response{Data: &gin.H{"id": id}}

		}
	}

	err := sr.repo.UpdateFilm(ctx, element.Title, element.Description, element.Language, element.ID, element.CategoryID)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: &gin.H{"id": element.ID}}
}

func (sr *AdminService) UpdateBook(ctx context.Context, form *multipart.Form, element models.ElementData, method string) models.Response {
	bookFilePath, bookImagePath, id := sr.repo.GetBookImageFilePath(ctx, element.ID)

	if id == 0 {
		return models.Response{Error: errors.New("not found"), Status: 404}
	}

	if method == "POST" {
		id := form.Value["id"][0]

		books := form.File["book"]
		images := form.File["image"]

		if len(books) > 1 || len(images) > 1 {
			return models.Response{Error: errors.New("too many files"), Status: 400}
		}

		if len(books) == 1 {

			if bookFilePath == "" {
				if bookImagePath != "" {
					bookFilePath = strings.TrimSuffix(bookImagePath, filepath.Ext(bookImagePath)) + ".pdf"
				} else {
					timestamp := time.Now().Unix()
					bookFilePath = fmt.Sprintf("/uploads/book/%d/%d.pdf", timestamp, timestamp)
				}
				sr.repo.UpdateBookPath(ctx, bookFilePath, element.ID)
			}

			bookExt := filepath.Ext(books[0].Filename)

			if bookExt != ".pdf" {
				return models.Response{Error: errors.New("invalid file type, must be  .pdf "), Status: 400}
			}

			os.RemoveAll("." + bookFilePath)
			time.Sleep(5 * time.Second)
			err := utils.SaveUploadedFile(books[0], "."+strings.TrimSuffix(bookFilePath, filepath.Ext(bookFilePath))+bookExt)

			if err != nil {
				return models.Response{Error: err, Status: 500}
			}

			// go utils.ConvertToHLS("."+strings.TrimSuffix(bookFilePath, filepath.Base(bookFilePath)), strings.TrimSuffix(filepath.Base(bookFilePath), filepath.Ext(filepath.Base(bookFilePath)))+bookExt, "film")

			return models.Response{Data: &gin.H{"id": id}}

		}

		if len(images) == 1 {

			if bookImagePath == "" {

				if bookFilePath != "" {
					bookImagePath = strings.TrimSuffix(bookFilePath, filepath.Ext(bookFilePath)) + ".webp"
				} else {
					timestamp := time.Now().Unix()
					bookImagePath = fmt.Sprintf("/uploads/book/%d/%d.webp", timestamp, timestamp)
				}

				sr.repo.UpdateBookImage(ctx, bookImagePath, element.ID)
			}

			allowedImageExts := map[string]bool{
				".jpg":  true,
				".jpeg": true,
				".png":  true,
				".webp": true,
			}

			imageEXT := filepath.Ext(images[0].Filename)
			if !allowedImageExts[imageEXT] {
				// todo: remove if uploaded film
				return models.Response{Error: errors.New("invalid file type, must be  .jpg, .jpeg, .png, .webp "), Status: 400}
			}

			os.RemoveAll("." + bookImagePath)
			err := utils.SaveUploadedFile(images[0], "."+strings.TrimSuffix(bookImagePath, filepath.Ext(bookImagePath))+imageEXT)

			if err != nil {
				return models.Response{Error: err, Status: 500}
			}

			fmt.Println("Ssss")

			status, err := utils.ResizeImage("."+strings.TrimSuffix(bookImagePath, filepath.Ext(bookImagePath))+imageEXT, 700)

			if err != nil {
				os.RemoveAll("." + bookImagePath)
				return models.Response{Error: err, Status: status}
			}

			return models.Response{Data: &gin.H{"id": id}}

		}
	}

	err := sr.repo.UpdateBook(ctx, element.Title, element.Description, element.Language, element.CategoryID, element.ID)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: &gin.H{"id": element.ID}}
}

func (sr *AdminService) UpdateMusic(ctx context.Context, form *multipart.Form, element models.ElementData, method string) models.Response {
	musicFilePath, musicImagePath, id := sr.repo.GetMusicImageFilePath(ctx, element.ID)

	if id == 0 {
		return models.Response{Error: errors.New("not found"), Status: 404}
	}

	if method == "POST" {
		id := form.Value["id"][0]

		musics := form.File["music"]
		images := form.File["image"]

		if len(musics) > 1 || len(images) > 1 {
			return models.Response{Error: errors.New("too many files"), Status: 400}
		}

		if len(musics) == 1 {

			if musicFilePath == "" {

				if musicImagePath != "" {
					musicFilePath = strings.TrimSuffix(musicImagePath, filepath.Base(musicImagePath)) + "/hls/" + strings.TrimSuffix(filepath.Base(musicImagePath), filepath.Ext(musicImagePath)) + ".m3u8"
				} else {
					timestamp := time.Now().Unix()
					musicFilePath = fmt.Sprintf("/uploads/music/%d/hls/%d.m3u8", timestamp, timestamp)
				}
				sr.repo.UpdateMusicPath(ctx, musicFilePath, element.ID)
			}

			musicExt := filepath.Ext(musics[0].Filename)

			if musicExt != ".mp3" {
				return models.Response{Error: errors.New("invalid file type, must be  .mp3 "), Status: 400}
			}

			os.RemoveAll("." + filepath.Dir(musicFilePath))
			time.Sleep(5 * time.Second)
			err := utils.SaveUploadedFile(musics[0], "."+strings.TrimSuffix(musicFilePath, filepath.Ext(musicFilePath))+musicExt)

			if err != nil {
				return models.Response{Error: err, Status: 500}
			}

			go utils.ConvertToHLS("."+strings.TrimSuffix(musicFilePath, filepath.Base(musicFilePath)), strings.TrimSuffix(filepath.Base(musicFilePath), filepath.Ext(filepath.Base(musicFilePath)))+musicExt, "music")

			return models.Response{Data: &gin.H{"id": id}}

		}

		if len(images) == 1 {

			if musicImagePath == "" {

				if musicFilePath != "" {
					startIndex := strings.Index(musicFilePath, "hls")
					musicImagePath = musicFilePath[:startIndex] + strings.TrimSuffix(filepath.Base(musicFilePath), filepath.Ext(musicFilePath)) + ".webp"
				} else {
					timestamp := time.Now().Unix()
					musicImagePath = fmt.Sprintf("/uploads/music/%d/%d.webp", timestamp, timestamp)
				}
				sr.repo.UpdateMusicImage(ctx, musicImagePath, element.ID)
			}

			allowedImageExts := map[string]bool{
				".jpg":  true,
				".jpeg": true,
				".png":  true,
				".webp": true,
			}

			imageEXT := filepath.Ext(images[0].Filename)
			if !allowedImageExts[imageEXT] {
				// todo: remove if uploaded music
				return models.Response{Error: errors.New("invalid file type, must be  .pdf and .jpg "), Status: 400}
			}

			os.RemoveAll("." + musicImagePath)
			err := utils.SaveUploadedFile(images[0], "."+strings.TrimSuffix(musicImagePath, filepath.Ext(musicImagePath))+imageEXT)

			if err != nil {
				return models.Response{Error: err, Status: 500}
			}

			fmt.Println("Ssss")

			status, err := utils.ResizeImage("."+strings.TrimSuffix(musicImagePath, filepath.Ext(musicImagePath))+imageEXT, 700)

			if err != nil {
				os.RemoveAll("." + musicImagePath)
				return models.Response{Error: err, Status: status}
			}

			return models.Response{Data: &gin.H{"id": id}}

		}
	}

	err := sr.repo.UpdateMusic(ctx, element.Title, element.Description, element.Language, element.ID, element.CategoryID)

	if err != nil {
		return models.Response{Error: err, Status: 500}
	}

	return models.Response{Data: &gin.H{"id": element.ID}}
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

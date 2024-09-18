package controllers

import (
	"media/internal/models"
	"media/internal/services"
	"media/pkg/utils"

	"github.com/YslamB/mglogger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminController struct {
	service *services.AdminService
	logger  *mglogger.Logger
}

func NewAdminController(db *pgxpool.Pool, logger *mglogger.Logger) *AdminController {
	return &AdminController{service: services.NewAdminService(db), logger: logger}
}

func (ctrl *AdminController) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := ctrl.service.GetUsers(ctx, 1)
	utils.GinResponse(c, 200, data, err, 0)
}

func (ctrl *AdminController) File(c *gin.Context) {

	ctx := c.Request.Context()
	form, err := c.MultipartForm()

	if err != nil {
		utils.GinResponse(c, 400, gin.H{"error": "Failed to parse multipart form"}, err, 0)
		return
	}

	data, err := ctrl.service.File(ctx, form)
	utils.GinResponse(c, 200, data, err, 0)
}

func (ctrl *AdminController) Music(c *gin.Context) {

	ctx := c.Request.Context()
	form, err := c.MultipartForm()

	if err != nil {
		utils.GinResponse(c, 400, gin.H{"error": "Failed to parse multipart form"}, err, 0)
		return
	}

	data, err := ctrl.service.Music(ctx, form)
	utils.GinResponse(c, 201, data, err, 0)
}

func (ctrl *AdminController) Film(c *gin.Context) {

	ctx := c.Request.Context()
	form, err := c.MultipartForm()

	if err != nil {
		utils.GinResponse(c, 400, gin.H{"error": "Failed to parse multipart form"}, err, 0)
		return
	}

	data, err := ctrl.service.Film(ctx, form)
	utils.GinResponse(c, 201, data, err, 0)
}

func (ctrl *AdminController) Book(c *gin.Context) {

	ctx := c.Request.Context()
	form, err := c.MultipartForm()

	if err != nil {
		utils.GinResponse(c, 400, gin.H{"error": "Failed to parse multipart form"}, err, 0)
		return
	}

	data, status, err := ctrl.service.Book(ctx, form)
	utils.GinResponse(c, 201, data, err, status)
}

func (ctrl *AdminController) DeleteMusic(c *gin.Context) {

	id := c.Param("id")
	ctx := c.Request.Context()
	err := ctrl.service.DeleteMusic(ctx, id)
	utils.GinResponse(c, 200, &gin.H{"message": "deleted"}, err, 404)

}

func (ctrl *AdminController) DeleteFilm(c *gin.Context) {

	id := c.Param("id")
	ctx := c.Request.Context()
	err := ctrl.service.DeleteFilm(ctx, id)
	utils.GinResponse(c, 200, &gin.H{"message": "deleted"}, err, 404)

}

func (ctrl *AdminController) DeleteBook(c *gin.Context) {

	id := c.Param("id")
	ctx := c.Request.Context()
	err := ctrl.service.DeleteBook(ctx, id)
	utils.GinResponse(c, 200, &gin.H{"message": "deleted"}, err, 404)

}

func (ctrl *AdminController) AdminLogin(c *gin.Context) {

	ctx := c.Request.Context()
	var admin models.LoginForm
	validationError := c.BindJSON(&admin)

	if validationError != nil {
		ctrl.logger.Println(validationError.Error())
		c.JSON(400, validationError.Error())
		return
	}

	acsessT, refT, err := ctrl.service.AdminLogin(ctx, admin)

	if err != nil {
		ctrl.logger.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{"access_token": acsessT, "refresh_token": refT})
}

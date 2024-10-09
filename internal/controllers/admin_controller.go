package controllers

import (
	"media/internal/models"
	"media/internal/services"
	"media/pkg/utils"
	"mime/multipart"
	"strconv"

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
	data := ctrl.service.GetUsers(ctx, 1)
	utils.GinResponse(c, data)
}

func (ctrl *AdminController) Music(c *gin.Context) {

	ctx := c.Request.Context()
	form, err := c.MultipartForm()

	if err != nil {
		utils.GinResponse(c, models.Response{Status: 400, Error: err})
		return
	}

	data := ctrl.service.Music(ctx, form)
	utils.GinResponse(c, data)
}

func (ctrl *AdminController) Film(c *gin.Context) {

	ctx := c.Request.Context()
	var reqBody models.ElementData
	validationError := c.BindJSON(&reqBody)

	if validationError != nil {
		utils.GinResponse(c, models.Response{Status: 400, Error: validationError})
		return
	}

	data := ctrl.service.Film(ctx, reqBody)
	utils.GinResponse(c, data)
}

func (ctrl *AdminController) Book(c *gin.Context) {

	ctx := c.Request.Context()
	var reqBody models.ElementData
	validationError := c.BindJSON(&reqBody)

	if validationError != nil {
		utils.GinResponse(c, models.Response{Status: 400, Error: validationError})
		return
	}

	data := ctrl.service.Book(ctx, reqBody)
	utils.GinResponse(c, data)
}

func (ctrl *AdminController) UpdateFilm(c *gin.Context) {

	var reqBody models.ElementData
	ctx := c.Request.Context()
	var form *multipart.Form

	if c.Request.Method == "POST" {
		s := c.PostForm("id")[0]
		reqBody.ID, _ = strconv.Atoi(string(s))

		form, _ = c.MultipartForm()
	} else {
		validationError := c.BindJSON(&reqBody)

		if validationError != nil {
			utils.GinResponse(c, models.Response{Status: 400, Error: validationError})
			return
		}
	}

	data := ctrl.service.UpdateFilm(ctx, form, reqBody, c.Request.Method)
	utils.GinResponse(c, data)
}

func (ctrl *AdminController) UpdateBook(c *gin.Context) {

	var reqBody models.ElementData
	ctx := c.Request.Context()
	var form *multipart.Form

	if c.Request.Method == "POST" {
		s := c.PostForm("id")[0]
		reqBody.ID, _ = strconv.Atoi(string(s))

		form, _ = c.MultipartForm()
	} else {
		validationError := c.BindJSON(&reqBody)

		if validationError != nil {
			utils.GinResponse(c, models.Response{Status: 400, Error: validationError})
			return
		}
	}

	data := ctrl.service.UpdateBook(ctx, form, reqBody, c.Request.Method)
	utils.GinResponse(c, data)
}

func (ctrl *AdminController) DeleteMusic(c *gin.Context) {

	id := c.Param("id")
	ctx := c.Request.Context()
	data := ctrl.service.DeleteMusic(ctx, id)
	utils.GinResponse(c, data)

}

func (ctrl *AdminController) DeleteFilm(c *gin.Context) {

	id := c.Param("id")
	ctx := c.Request.Context()
	data := ctrl.service.DeleteFilm(ctx, id)
	utils.GinResponse(c, data)

}

func (ctrl *AdminController) DeleteBook(c *gin.Context) {

	id := c.Param("id")
	ctx := c.Request.Context()
	data := ctrl.service.DeleteBook(ctx, id)
	utils.GinResponse(c, data)

}

func (ctrl *AdminController) AdminLogin(c *gin.Context) {

	ctx := c.Request.Context()
	var admin models.LoginForm
	validationError := c.BindJSON(&admin)

	if validationError != nil {
		utils.GinResponse(c, models.Response{Status: 400, Error: validationError})
		return
	}

	acsessT, refT, err := ctrl.service.AdminLogin(ctx, admin)

	if err != nil {
		utils.GinResponse(c, models.Response{Status: 500, Error: err})
		return
	}

	utils.GinResponse(c, models.Response{Data: gin.H{"access_token": acsessT, "refresh_token": refT}})
}

func (ctrl *AdminController) Category(c *gin.Context) {

	ctx := c.Request.Context()
	var category models.Category
	validationError := c.BindJSON(&category)

	if validationError != nil {
		utils.GinResponse(c, models.Response{Status: 400, Error: validationError})
		return
	}

	data := ctrl.service.Category(ctx, category)
	utils.GinResponse(c, data)

}

func (ctrl *AdminController) SubCategory(c *gin.Context) {

	ctx := c.Request.Context()
	var category models.Category
	validationError := c.BindJSON(&category)

	if validationError != nil {
		utils.GinResponse(c, models.Response{Status: 400, Error: validationError})
		return
	}

	data := ctrl.service.SubCategory(ctx, category)
	utils.GinResponse(c, data)

}

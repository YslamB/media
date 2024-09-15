package controllers

import (
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

	utils.GinResponse(c, 200, data, err, 0)

}

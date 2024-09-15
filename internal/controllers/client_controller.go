package controllers

import (
	"media/internal/services"
	"media/pkg/utils"

	"github.com/YslamB/mglogger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientController struct {
	service *services.ClientService
	logger  *mglogger.Logger
}

func NewClientController(db *pgxpool.Pool, logger *mglogger.Logger) *ClientController {
	return &ClientController{service: services.NewClientService(db), logger: logger}
}

func (u *ClientController) Films(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := u.service.GetUsers(ctx, 1)
	utils.GinResponse(c, 200, data, err, 0)
}

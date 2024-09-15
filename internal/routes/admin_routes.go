package routes

import (
	"media/internal/controllers"

	"github.com/gin-gonic/gin"
)

func setupAdminRoutes(r *gin.RouterGroup, ctrl *controllers.AdminController) {

	r.GET("", ctrl.GetUsers)
	r.POST("/music", ctrl.Music)

}

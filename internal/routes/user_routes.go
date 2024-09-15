package routes

import (
	"media/internal/controllers"
	"media/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func setupClientRoutes(r *gin.RouterGroup, ctrl *controllers.ClientController) {
	r.GET("/films", middlewares.AdminGuard, middlewares.PageLimitSet, ctrl.Films)
	// r.POST("/buy-sub", middlewares.UserGuard, ctrl.BuySubscription)
	// r.PUT("", middlewares.UserGuard, ctrl.UpdateUser)

	// // 				TODO: not delete, just disable user and delete order's files.
	// r.DELETE("/:id", middlewares.UserGuard, middlewares.ParamIDToInt, ctrl.DeleteUser)
}

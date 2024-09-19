package routes

import (
	"media/internal/controllers"
	"media/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func setupClientRoutes(r *gin.RouterGroup, ctrl *controllers.ClientController) {

	r.GET("/films", middlewares.PageLimitSet, ctrl.Films)
	r.GET("/musics", middlewares.PageLimitSet, ctrl.Books)
	r.GET("/books", middlewares.PageLimitSet, ctrl.Musics)

}

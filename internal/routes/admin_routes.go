package routes

import (
	"media/internal/controllers"

	"github.com/gin-gonic/gin"
)

func setupAdminRoutes(r *gin.RouterGroup, ctrl *controllers.AdminController) {

	r.GET("", ctrl.GetUsers)
	r.POST("/login", ctrl.AdminLogin)
	r.POST("/music", ctrl.Music)
	r.POST("/film", ctrl.Film)
	r.POST("/book", ctrl.Book)

	r.DELETE("/music/:id", ctrl.DeleteMusic)
	r.DELETE("/film/:id", ctrl.DeleteFilm)
	r.DELETE("/book/:id", ctrl.DeleteBook)

}

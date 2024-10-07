package routes

import (
	"media/internal/controllers"
	"media/pkg/middlewares"

	"github.com/YslamB/mglogger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(r *gin.Engine, db *pgxpool.Pool, logger *mglogger.Logger) {

	rClient := r.Group("media/client")
	clientC := controllers.NewClientController(db, logger)

	{
		rClient.GET("/categories", middlewares.PageLimitSet, clientC.Categories)
		rClient.GET("/films", middlewares.PageLimitSet, clientC.Films)
		rClient.GET("/books", middlewares.PageLimitSet, clientC.Books)
		rClient.GET("/musics", middlewares.PageLimitSet, clientC.Musics)
	}

	rAdmin := r.Group("media/admin", middlewares.Guard)
	adminC := controllers.NewAdminController(db, logger)

	{
		rAdmin.GET("", middlewares.AdminGuard, adminC.GetUsers)
		rAdmin.POST("/category", middlewares.AdminGuard, adminC.Category)
		rAdmin.POST("/sub-category", middlewares.AdminGuard, adminC.SubCategory)
		rAdmin.POST("/music", middlewares.AdminGuard, adminC.Music)
		rAdmin.POST("/film", middlewares.AdminGuard, adminC.Film)
		rAdmin.POST("/book", middlewares.AdminGuard, adminC.Book)
		rAdmin.POST("/login", adminC.AdminLogin)

		rAdmin.DELETE("/music/:id", middlewares.AdminGuard, adminC.DeleteMusic)
		rAdmin.DELETE("/film/:id", middlewares.AdminGuard, adminC.DeleteFilm)
		rAdmin.DELETE("/book/:id", middlewares.AdminGuard, adminC.DeleteBook)
	}

}

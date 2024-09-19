package routes

import (
	"media/internal/controllers"

	"github.com/YslamB/mglogger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(r *gin.Engine, db *pgxpool.Pool, logger *mglogger.Logger) {

	rClient := r.Group("media/client")
	clientController := controllers.NewClientController(db, logger)
	setupClientRoutes(rClient, clientController)

	rAdmin := r.Group("media/admin")
	adminController := controllers.NewAdminController(db, logger)
	setupAdminRoutes(rAdmin, adminController)

}

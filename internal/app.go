package src

import (
	"media/internal/routes"
	"media/pkg/config"
	"media/pkg/middlewares"
	"media/pkg/utils"
	"os"

	"github.com/YslamB/mglogger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitApp(db *pgxpool.Pool, logger *mglogger.Logger) *gin.Engine {
	// rl := middlewares.NewRateLimiter()

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	router := gin.New()
	router.SetTrustedProxies(nil)

	// Apply the rate limiter middleware to all routes.
	// router.Use(middlewares.RateLimiterMiddleware(rl))
	router.Use(middlewares.Cors)
	router.Use(utils.MGLogger(logger))

	router.Static("/media/uploads", config.ENV.UPLOAD_PATH)

	// new routers
	routes.SetupRoutes(router, db, logger)
	return router
}

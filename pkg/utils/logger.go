package utils

import (
	"media/pkg/config"
	"time"

	"github.com/YslamB/mglogger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Log *mglogger.Logger

func InitLogger() *mglogger.Logger {
	Log = mglogger.GetLogger(config.ENV.LOGGER_FOLDER_PATH, config.ENV.LOGGER_FILENAME, config.ENV.GIN_MODE)
	return Log
}

func MGLogger(logger *mglogger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(startTime)

		// Get the status code of the response
		statusCode := c.Writer.Status()

		// Log the request details using logrus
		logger.WithFields(logrus.Fields{
			"status_code": statusCode,
			"latency":     latency,
			"client_ip":   c.ClientIP(),
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"user_agent":  c.Request.UserAgent(),
			"error":       c.Errors.ByType(gin.ErrorTypePrivate).String(),
		}).Info("Request details")
	}
}

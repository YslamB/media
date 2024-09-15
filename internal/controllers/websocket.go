package controllers

import (
	"media/internal/services"

	"github.com/gin-gonic/gin"
)

func WebSocket(router *gin.Engine) {
	group := router.Group("ussa/ws")
	group.GET("/:token", services.HandleConnections)
}

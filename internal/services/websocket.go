package services

import (
	"media/pkg/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

func HandleConnections(c *gin.Context) {
	token := c.Param("token")

	if token == "" {
		c.AbortWithStatus(401)
		return
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		token, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.ENV.ACCESS_KEY), nil
		},
	)

	if err != nil {
		config.SocketClients[token] = nil
		c.AbortWithStatus(403)
		return
	}

	role := claims["role"].(string)
	userID := int(claims["id"].(float64))
	user := role + " " + strconv.Itoa(userID)

	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

	ws, _ := upgrader.Upgrade(c.Writer, c.Request, nil)
	config.SocketClients[user] = ws
}

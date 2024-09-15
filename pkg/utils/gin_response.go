package utils

import (
	"github.com/gin-gonic/gin"
)

func GinResponse(c *gin.Context, status int, data any, err error, errorStatus int) {

	if err != nil {
		Log.Println(err.Error())

		if errorStatus != 0 {
			c.JSON(errorStatus, err.Error())
			return
		}

		c.JSON(500, err.Error())
		return
	}

	c.JSON(status, data)
}

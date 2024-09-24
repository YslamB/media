package middlewares

import (
	"media/pkg/config"
	"media/pkg/utils"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/time/rate"
)

func Guard(c *gin.Context) {
	authorization := c.Request.Header["Authorization"]

	if len(authorization) == 0 {
		c.AbortWithStatus(401)
		return
	}

	bearer := strings.Split(authorization[0], "Bearer ")

	if len(bearer) == 0 || len(bearer) == 1 {
		c.AbortWithStatus(401)
		return
	}

	token := bearer[1]
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(
		token, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.ENV.ACCESS_KEY), nil
		},
	)

	if err != nil {
		utils.Log.Println(err.Error())
		c.AbortWithStatus(403)
		return
	}

	c.Set("id", claims["id"].(string))
	c.Set("role", claims["role"])

	c.Next()
}

func AdminGuard(c *gin.Context) {
	role := c.MustGet("role").(string)

	if role != "admin" {
		c.AbortWithStatus(403)
		return
	}

	c.Next()
}

func ParamIDToInt(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		c.AbortWithStatus(400)
		return
	}
	id, err := strconv.Atoi(idStr)

	if err != nil || id <= 0 {
		c.JSON(400, gin.H{"message": "ID must be positive integer number"})
		return
	}

	c.Set("paramID", id)
	c.Next()
}

func PageLimitSet(c *gin.Context) {
	pageStr := c.Query("page")
	countStr := c.Query("count")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(countStr)
	if err != nil {
		limit = 20
	}

	c.Set("page", page)
	c.Set("limit", limit)
	c.Next()
}

type RateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.RWMutex
}

// NewRateLimiter creates a new RateLimiter instance.
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limiters: make(map[string]*rate.Limiter),
	}
}

// GetLimiter returns a rate limiter for the given device ID.
func (rl *RateLimiter) getLimiter(deviceID string) *rate.Limiter {
	rl.mu.RLock()
	limiter, exists := rl.limiters[deviceID]
	rl.mu.RUnlock()

	if !exists {
		rl.mu.Lock()
		// Create a new rate limiter with 1 request per second and a burst of 3 requests.
		limiter = rate.NewLimiter(1, 3)
		rl.limiters[deviceID] = limiter
		rl.mu.Unlock()
	}

	return limiter
}

// RateLimiterMiddleware creates a middleware that applies rate limiting based on device ID.
func RateLimiterMiddleware(rl *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		deviceID := c.GetHeader("X-Header-Device-Id")
		if deviceID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing X-Header-Device-Id"})
			c.Abort()
			return
		}

		limiter := rl.getLimiter(deviceID)

		// Check if the rate limiter allows the request
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
			c.Abort()
			return
		}

		c.Next()
	}
}

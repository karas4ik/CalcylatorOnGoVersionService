package handler

import (
	"calc_service/internal/cache"
	"calc_service/internal/logger"
	"calc_service/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.GetLogger()
		log.Infof("Request: %s %s", c.Request.Method, c.Request.URL)
		c.Next()
		log.Infof("Response: %d", c.Writer.Status())
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log := logger.GetLogger()
				log.Errorf("Recovered from panic: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}()
		c.Next()
	}
}

func CacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		expression := c.PostForm("expression")
		if result, found := cache.Get(expression); found {
			c.JSON(http.StatusOK, model.Response{Result: result})
			return
		}

		c.Next()

		if c.Writer.Status() == http.StatusOK {
			cache.Set(expression, c.Writer.String())
		}
	}
}

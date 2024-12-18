package r

import (
	"calc_service/config"
	"calc_service/internal/handler"
	"calc_service/internal/logger"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() error {
	cfg := config.LoadConfig()
	logger.InitLogger(cfg.LogLevel)
	log := logger.GetLogger()
	log.Info("Starting the calculator service...")

	r := gin.Default()
	r.Use(handler.LoggingMiddleware())
	r.Use(handler.RecoveryMiddleware())
	r.Use(handler.CacheMiddleware())
	r.POST("/api/v1/calculate", handler.CalculateHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method Not Allowed"})
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err1 := server.ListenAndServe(); err1 != nil && err1 != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
		return err
	}
	log.Info("Server exiting")
	return nil
}

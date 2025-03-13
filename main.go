package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	sdkginext "github.com/larksuite/oapi-sdk-gin"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	"go.uber.org/zap"

	"feishuBot/internal/conf"
	"feishuBot/internal/handlers"
	"feishuBot/internal/services"
	"feishuBot/utils/logger"
)

func main() {
	// Initialize configuration
	if err := conf.InitConf(); err != nil {
		logger.Fatal("Failed to initialize configuration", zap.Error(err))
	}
	logger.Info("Configuration initialized successfully")

	// Initialize services
	services.InitLark()
	services.InitOpenAI()
	logger.Info("Services initialized successfully")

	// Set up Feishu event handler
	handler := dispatcher.NewEventDispatcher(conf.GConfig.Lark.VerifyToken, "")
	handler = handler.OnP2MessageReceiveV1(handlers.ReceiveMsgHandler)

	// Set up gin router
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()

	api := g.Group("/api/v1")
	{
		api.POST("/feishu/webhook", sdkginext.NewEventHandlerFunc(handler))
	}

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":8081",
		Handler: g,
	}

	// Start server in a goroutine
	go func() {
		logger.Info(fmt.Sprintf("Server started successfully, listening on port: %s", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// Set shutdown timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shutdown server
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server shutdown error", zap.Error(err))
	}

	logger.Info("Server shutdown completed")
}

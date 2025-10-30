package main

import (
	"context"
	"errors"
	"fmt"
	"gin-basic/config"
	"gin-basic/utils/logger"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.GetConfig()

	r := gin.Default()

	logger.GetLogger()

	r.GET("/health", func(ctx *gin.Context) {
		logger.Log.Infoln("GET /health")
		ctx.JSON(http.StatusOK, "ok")
	})

	port := "8080"
	if cfg.App.Api.Port != "" {
		port = cfg.App.Api.Port
	}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Fatalf("[API服务器] [故障] [Message]:错误信息: %s\n", err)
		} else {
			logger.Log.Infof("[API服务器] [Message]:服务器已启动,监听端口:%s", port)
		}
	}()

	<-ctx.Done()

	stop()
	logger.Log.Infoln("[API服务器] [Message]:正在关闭...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatalf("[API服务器] 关闭失败... %v", err)
	}

	logger.Log.Infoln("[API服务器] [Message]:服务器已退出")
}

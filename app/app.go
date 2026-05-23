package app

import (
	"context"
	"errors"
	"fmt"
	"gin-basic/internal/bootstrap"
	"gin-basic/utils/logger"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 初始化启动容器, 包含配置、工具、服务、控制器
	c := bootstrap.NewContainer()

	app := bootstrap.NewRouter(c)

	port := "8080"
	if c.Config.App.Api.Port != "" {
		port = c.Config.App.Api.Port
	}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app,
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

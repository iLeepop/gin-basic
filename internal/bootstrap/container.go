package bootstrap

import (
	"gin-basic/internal/cfg"
	"gin-basic/internal/ports/iservice"
	"gin-basic/internal/utils"
	"gin-basic/internal/web"
	"gin-basic/utils/logger"
)

type Container struct {
	// 配置
	Config *cfg.Configuration

	// 工具
	JwtUtils *utils.JwtUtils

	// 服务
	ServerService iservice.IServerService

	// 控制器
	ServerController *web.ServerController
}

func NewContainer() *Container {
	c := &Container{}

	c.Init()

	buildConfig(c)
	buildUtils(c)
	buildService(c)
	buildController(c)

	return c
}

func (c *Container) Init() {
	// 初始化 logger
	logger.InitLogger(c.Config)
}

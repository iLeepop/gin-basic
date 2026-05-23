package bootstrap

import (
	"gin-basic/config"
	server_impl "gin-basic/internal/service"
	"gin-basic/internal/utils"
	"gin-basic/internal/web"
	"gin-basic/utils/logger"
)

type Container struct {
	// 配置
	Config *config.Configuration

	// 工具
	JwtUtils *utils.JwtUtils

	// 服务
	ServerService *server_impl.ServerService

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
	// 初始化配置
	config.InitConfig()

	// 初始化 logger
	logger.InitLogger()
}

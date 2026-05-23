package bootstrap

import (
	"gin-basic/config"
	"gin-basic/internal/cfg"
)

func buildConfig(c *Container) {
	cfg := config.GetConfig[cfg.Configuration]()
	c.Config = cfg.Config
}

package bootstrap

import "gin-basic/config"

func buildConfig(c *Container) {
	c.Config = config.GetConfig()
}

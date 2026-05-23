package bootstrap

import "gin-basic/internal/web"

func buildController(c *Container) {
	c.ServerController = &web.ServerController{
		ServerService: c.ServerService,
	}
}

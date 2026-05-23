package bootstrap

import server "gin-basic/internal/service"

func buildService(c *Container) {
	c.ServerService = &server.ServerService{}
}

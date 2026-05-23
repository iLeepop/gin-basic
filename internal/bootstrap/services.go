package bootstrap

import server_impl "gin-basic/internal/service"

func buildService(c *Container) {
	c.ServerService = &server_impl.ServerService{}
}

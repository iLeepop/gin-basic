package web

import (
	"gin-basic/internal/core"
	"gin-basic/internal/ports/service"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
	ServerService service.IServerService
}

func (c *ServerController) HealthCheck(ctx *gin.Context) {
	result, err := c.ServerService.HealthCheck()
	if err != nil {
		core.ToResult(ctx, nil, err)
		return
	}
	core.ToResult(ctx, result, nil)
}

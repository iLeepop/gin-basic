package bootstrap

import (
	"gin-basic/internal/utils"
	"time"
)

func buildUtils(c *Container) {
	c.JwtUtils = &utils.JwtUtils{
		Secret:   c.Config.Jwt.Secret,
		TokenTTL: 1 * time.Hour,
	}
}

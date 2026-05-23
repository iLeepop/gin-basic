package bootstrap

import (
	"gin-basic/internal/utils"
	"time"
)

func buildUtils(c *Container) {
	c.JwtUtils = &utils.JwtUtils{
		Secret:   "custom_secret",
		TokenTTL: 1 * time.Hour,
	}
}

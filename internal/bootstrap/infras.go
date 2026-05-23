package bootstrap

import (
	"gin-basic/internal/infras/cache"
	"gin-basic/internal/infras/db"
	"gin-basic/utils/logger"
)

func buildInfras(c *Container) {
	mysql, err := db.NewMySQL(c.Config.Database.Mysql, c.Config.Database.ConnectPool)
	if err != nil {
		logger.Log.Fatalf("[基础设施] [MySQL] 初始化失败: %v", err)
	}
	c.MySQL = mysql

	postgresql, err := db.NewPostgreSQL(c.Config.Database.PostgreSQL, c.Config.Database.ConnectPool)
	if err != nil {
		logger.Log.Fatalf("[基础设施] [PostgreSQL] 初始化失败: %v", err)
	}
	c.PostgreSQL = postgresql

	redis, err := cache.NewRedis(c.Config.Cache.Redis)
	if err != nil {
		logger.Log.Fatalf("[基础设施] [Redis] 初始化失败: %v", err)
	}
	c.Redis = redis
}

func (c *Container) CloseInfras() {
	if c.MySQL != nil {
		_ = c.MySQL.Close()
	}
	if c.PostgreSQL != nil {
		_ = c.PostgreSQL.Close()
	}
	if c.Redis != nil {
		_ = c.Redis.Close()
	}
}

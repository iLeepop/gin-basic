package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// 配置
type Config[T any] struct {
	Config *T
}

var _v *viper.Viper

func GetConfig[T any]() *Config[T] {
	_cfg := &Config[T]{}

	// 读取 config.yaml中文件
	_v = viper.New()
	_v.SetConfigName("config")
	_v.SetConfigType("yml")
	_v.AddConfigPath("./conf")

	if err := _v.ReadInConfig(); err != nil {
		panic(err)
	}

	// 开发阶段使用.env文件
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println("没有本地.env文件")
	}

	// 使用环境变量
	_v.AutomaticEnv()
	//_v.SetEnvPrefix("JYW")
	_v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 将配置文件映射到结构体中
	if err := _v.Unmarshal(&_cfg.Config); err != nil {
		panic(err)
	}

	return _cfg
}

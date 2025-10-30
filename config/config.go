package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Configuration struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Log   Log   `mapstructure:"log" json:"log" yaml:"log"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Name    string `mapstructure:"name" json:"name" yaml:"name"`
	Version string `mapstructure:"version" json:"version" yaml:"version"`
	Api     Server `mapstructure:"api" json:"api" yaml:"api"`
}

type Server struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}

type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	Output     string `mapstructure:"output" json:"output" yaml:"output"`
	FilePath   string `mapstructure:"file_path" json:"file_path" yaml:"file_path"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type Mysql struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

// var _config Config
var _config Configuration
var _v *viper.Viper

func init() {
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
	_config = Configuration{}
	if err := _v.Unmarshal(&_config); err != nil {
		panic(err)
	}

	if _config.App.Env == "dev" {
		fmt.Println("[配置] [初始化] [信息]:", _config)
	}
}

//func GetConfig() Config {
//	return _config
//}

func GetConfig() Configuration {
	return _config
}

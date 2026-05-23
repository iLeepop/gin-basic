package cfg

// 配置结构体
type Configuration struct {
	App      App      `mapstructure:"app" json:"app" yaml:"app"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Cache    Cache    `mapstructure:"cache" json:"cache" yaml:"cache"`
	Jwt      Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

// 应用配置
type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Name    string `mapstructure:"name" json:"name" yaml:"name"`
	Version string `mapstructure:"version" json:"version" yaml:"version"`
	Api     Server `mapstructure:"api" json:"api" yaml:"api"`
}

// 服务配置
type Server struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}

// 日志配置
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

// 数据库配置
type Database struct {
	ConnectPool ConnectPool `mapstructure:"connect_pool" json:"connect_pool" yaml:"connect_pool"`
	Mysql       Mysql       `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PostgreSQL  PostgreSQL  `mapstructure:"postgresql" json:"postgresql" yaml:"postgresql"`
}

// 连接池配置
type ConnectPool struct {
	MaxOpenConns    int `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	MaxIdleConns    int `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	ConnMaxLifetime int `mapstructure:"conn_max_lifetime" json:"conn_max_lifetime" yaml:"conn_max_lifetime"`
}

// 缓存配置
type Cache struct {
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}

// mysql 配置
type Mysql struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
}

// postgresql 配置
type PostgreSQL struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
}

// redis 配置
type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

// jwt 配置
type Jwt struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
}

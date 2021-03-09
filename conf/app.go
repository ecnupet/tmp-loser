package conf

type AppConfig struct {
	AppName     string       `toml:"app_name" validate:"required"`
	ServerHost  string       `toml:"server_host " validate:"required"`
	GinMode     string       `toml:"gin_mode" validate:"required"`
	MysqlConfig *MysqlConfig `toml:"mysql" validate:"required"`
	RedisConfig *RedisConfig `toml:"redis" validate:"required"`
}
type MysqlConfig struct {
	Dsn          string `toml:"dsn" validate:"required"`
	MaxIdleConns int `toml:"max_idle_conns" validate:"required"`
	MaxOpenConns int `toml:"max_open_conns" validate:"required"`
}

type RedisConfig struct {
	Addr     string `toml:"addr" validate:"required"`
	PoolSize int    `toml:"pool_size" validate:"required"`
}

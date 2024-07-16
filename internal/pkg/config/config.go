package config

var (
	cfgInstance *Config
)

type Config struct {
	Environment string `env:"ENV,required"`

	DatabaseURL     string `env:"DATABASE_URL,required"`
	DatabaseMaxConn int    `env:"DATABASE_MAX_CONN" envDefault:"100"`
	DatabaseMaxIdle int    `env:"DATABASE_MAX_IDLE" envDefault:"10"`
	DatabaseTimeout int    `env:"DATABASE_TIMEOUT" envDefault:"30"`

	HttpPort int `env:"HTTP_PORT" envDefault:"8080"`
}

func GetConfig() *Config {
	return cfgInstance
}

type Initializer func() (*Config, error)

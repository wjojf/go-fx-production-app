package config

var (
	cfgInstance *Config
)

type Initializer func() (*Config, error)

type Config struct {
	Environment string `env:"ENV,required"`

	DatabaseURL     string `env:"DATABASE_URL,required"`
	DatabaseMaxConn int    `env:"DATABASE_MAX_CONN" envDefault:"100"`
	DatabaseMaxIdle int    `env:"DATABASE_MAX_IDLE" envDefault:"10"`
	DatabaseTimeout int    `env:"DATABASE_TIMEOUT" envDefault:"30"`

	HttpPort int `env:"HTTP_PORT" envDefault:"8080"`

	GoogleProjectID string `env:"GOOGLE_PROJECT_ID,required"`

	JaegerUrl string `env:"JAEGER_URL"`

	JwtSigningKey                string `env:"JWT_SIGN_KEY,required"`
	JwtAccessTokenLifetimeHours  int    `env:"JWT_ACCESS_TOKEN_LIFETIME_HOURS" envDefault:"1"`
	JwtRefreshTokenLifetimeHours int    `env:"JWT_REFRESH_TOKEN_LIFETIME_HOURS" envDefault:"24"`
}

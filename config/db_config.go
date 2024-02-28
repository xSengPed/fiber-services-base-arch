package config

type DbConfig struct {
	Host     string `env:"DB_HOST,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	DbName   string `env:"DB_NAME,required"`
	Port     string `env:"DB_PORT,required"`
	Ssl      string `env:"DB_SSL,required"`
	Timezone string `env:"DB_TIMEZONE,required"`
}

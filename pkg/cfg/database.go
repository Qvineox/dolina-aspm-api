package cfg

type DatabaseConfig struct {
	Host string `json:"host" yaml:"host" env:"HOST" env-default:"0.0.0.0" validate:"required"`
	Port uint32 `json:"port" yaml:"port" env:"PORT" env-default:"5432" validate:"required"`

	User string `json:"user" yaml:"user" env:"USER" env-required:"true" validate:"required"`
	Pass string `json:"pass" yaml:"pass" env:"PASS"`

	Timezone string `json:"timezone" yaml:"timezone" env:"TIMEZONE" env-default:"Europe/Moscow"`
}

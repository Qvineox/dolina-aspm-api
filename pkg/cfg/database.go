package cfg

import (
	"fmt"
)

type DatabaseConfig struct {
	Host string `json:"host" yaml:"host" env:"HOST" env-default:"0.0.0.0" validate:"required"`
	Port uint32 `json:"port" yaml:"port" env:"PORT" env-default:"5432" validate:"required"`
	Name string `json:"name" yaml:"name" env:"NAME"`

	User string `json:"user" yaml:"user" env:"USER" env-required:"true" validate:"required"`
	Pass string `json:"pass" yaml:"pass" env:"PASS"`

	Timezone string `json:"timezone" yaml:"timezone" env:"TIMEZONE" env-default:"Europe/Moscow"`

	TraceAllMessages bool `json:"trace_all_messages" yaml:"trace_all_messages" env:"TRACE_ALL"`
}

type DatabaseLogsConfig struct {
}

func (config DatabaseConfig) Postgres() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable TimeZone=%s",
		config.Host,
		config.Port,
		config.User,
		config.Pass,
		config.Name,
		config.Timezone,
	)
}

package cfg

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

var configValidator = validator.New()

type StaticConfig struct {
	Database DatabaseConfig `json:"database" yaml:"database" env-prefix:"DB_"`
	Logging  LoggingConfig  `json:"logging" yaml:"logging" env-prefix:"LOG_"`
}

func NewConfigFromFile(path string) (StaticConfig, error) {
	var sc StaticConfig

	err := cleanenv.ReadConfig(path, &sc)
	if err != nil {
		return StaticConfig{}, fmt.Errorf("error reading config file: %w", err)
	}

	err = configValidator.Struct(sc)
	if err != nil {
		return StaticConfig{}, fmt.Errorf("failed to validate config: %w", err)
	}

	return sc, nil
}

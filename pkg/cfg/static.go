package cfg

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type StaticConfig struct {
	Database DatabaseConfig `json:"database" yaml:"database" env-prefix:"DB_"`
}

func (sc *StaticConfig) Validate() error {
	validator.

	return nil
}

func NewConfigFromFile(path string) StaticConfig {
	var sc StaticConfig

	err := cleanenv.ReadConfig(path, &sc)
	if err != nil {
		panic(fmt.Errorf("error reading config file: %w", err))
	}

	return sc
}

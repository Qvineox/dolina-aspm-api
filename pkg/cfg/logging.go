package cfg

type LoggingConfig struct {
	Level int `json:"level" yaml:"level" env:"LEVEL" env-default:"0" validate:"min=-4,max=8"`
}

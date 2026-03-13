package cfg

type LogsConfig struct {
	Host        string `json:"host" yaml:"host"`
	Environment string `json:"environment" yaml:"environment" env-default:"dev" validate:"required,oneof=dev test stage prod"`

	Level     int  `json:"level" yaml:"level" env:"LEVEL" env-default:"0" validate:"min=-4,max=8"`
	AddSource bool `json:"add_source" yaml:"add_source"`

	Format   string `json:"format" yaml:"format" env-default:"text" validate:"required,oneof=json text"`
	Filepath string `json:"filepath" yaml:"filepath"`
}

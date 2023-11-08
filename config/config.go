package config

type Config struct {
	Mysql Mysql `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	Redis Redis `json:"redis" yaml:"redis" mapstructure:"redis"`
	Zap   Zap   `json:"zap" yaml:"zap" mapstructure:"zap"`
	Jwt   Jwt   `json:"jwt" yaml:"jwt" mapstructure:"jwt"`
}

package config

type Redis struct {
	DB       int    `json:"db" yaml:"db" mapstructure:"db"`
	Addr     string `json:"addr" yaml:"addr" mapstructure:"addr"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
}

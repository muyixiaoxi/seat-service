package config

type Jwt struct {
	SigningKey string `json:"signing-key" yaml:"signing-key" mapstructure:"signing-key"`
	Expires    int    `json:"expires" yaml:"expires" mapstructure:"expires"`
	Issuer     string `json:"issuer" yaml:"issuer" mapstructure:"issuer"`
	Buffer     int    `json:"buffer" yaml:"buffer" mapstructure:"buffer"`
}

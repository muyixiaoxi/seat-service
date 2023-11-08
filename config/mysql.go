package config

type Mysql struct {
	Path         string `json:"path" yaml:"path" mapstructure:"path"`
	Port         string `json:"port" yaml:"port" mapstructure:"port"`
	Config       string `json:"config" yaml:"config" mapstructure:"config"`
	DBName       string `json:"db-name" yaml:"db-name" mapstructure:"db-name"`
	Username     string `json:"username" yaml:"username" mapstructure:"username"`
	Password     string `json:"password" yaml:"password" mapstructure:"password"`
	MaxIdleConns int    `json:"max-idle-conns" yaml:"max-idle-conns" mapstructure:"max-idle-conns"`
	MaxOpenConns int    `json:"max-open-conns" yaml:"max-open-conns" mapstructure:"max-open-conns"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DBName + "?" + m.Config
}

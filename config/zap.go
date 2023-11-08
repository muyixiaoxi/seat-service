package config

type Zap struct {
	Level         string `json:"level" yaml:"level" mapstructure:"level" `
	Format        string `json:"format" yaml:"format" mapstructure:"format" `
	Prefix        string `json:"prefix" yaml:"prefix" mapstructure:"prefix" `
	Director      string `json:"director" yaml:"director" mapstructure:"director" `
	ShowLine      bool   `json:"show-line" yaml:"show-line" mapstructure:"show-line" `
	EncodeLevel   string `json:"encode-level" yaml:"encode-level" mapstructure:"encode-level" `
	StacktraceKey string `json:"stacktrace-key" yaml:"stacktrace-key" mapstructure:"stacktrace-key" `
	LogInConsole  bool   `json:"log-in-console" yaml:"log-in-console" mapstructure:"log-in-console" `
	FilePath      string `json:"file-path" yaml:"file-path" mapstructure:"file-path"`
}

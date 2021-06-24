package api

type Config struct {
	// application port
	BindAddr string `toml:"bind_addr"`

	// application log level
	LoggerLevel string `toml:"logger_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "info",
	}
}

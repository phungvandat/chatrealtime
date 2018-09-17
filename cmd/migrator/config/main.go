package config

// Config contain configuration of db for migrator
type Config struct {
	DBType          string
	DBUserName      string
	DBPassword      string
	DBName          string
	DBSSLModeOption string
	DBHostname      string
	DBPort          string
	DBEnvironment   string
}

// Reader get config from reader
type Reader interface {
	Read() (*Config, error)
}

// GetBy get config by reader(file or env)
func GetBy(r Reader) (*Config, error) {
	return &Config{
		DBType:          "postgres",
		DBUserName:      "postgres",
		DBPassword:      "chatsexeverywhere",
		DBName:          "chat-sex",
		DBSSLModeOption: "disable",
		DBHostname:      "localhost",
		DBPort:          "5432",
	}, nil
}

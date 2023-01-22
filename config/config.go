package config

import "os"

type Config struct {
	Port    string
	Host    string
	User    string
	Pass    string
	DBName  string
	SSLMode string
}

func GetConfig() *Config {
	config := &Config{}
	config.Port = os.Getenv("DB_PORT")
	config.Host = os.Getenv("DB_HOST")
	config.User = os.Getenv("DB_USER")
	config.Pass = os.Getenv("DB_PASS")
	config.DBName = os.Getenv("DB_DBNAME")
	config.SSLMode = os.Getenv("DB_SSL_MODE")
	return config
}

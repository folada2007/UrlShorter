package apiserver

import "fmt"

type DBConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"dbname"`
	SSLMode  string `toml:"sslmode"`
}

type Config struct {
	BindAddr string   `toml:"bind_addr"`
	LogLevel string   `toml:"log_level"`
	Db       DBConfig `toml:"db"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Db: DBConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			Database: "postgres",
			SSLMode:  "disable",
		},
	}
}

func (db *DBConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s, port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.Database, db.SSLMode)
}

package config

import (
	"os"
)

var (
	Env *AppConfig
)

const (
	keyEnv               = "ENV"
	keyHost              = "HOST"
	keyPort              = "PORT"
	keyAccessAllowOrigin = "ACCESS_ALLOW_ORIGIN"
	keyDatabaseDSN       = "DATABASE_DSN"
	keyLogFolderPath     = "LOG_FOLDER_PATH"
)

type AppConfig struct {
	Env               string
	Host              string
	Port              string
	AccessAllowOrigin string
	DatabaseDSN       string
	DataPath          *DataPathConfig
}

type DataPathConfig struct {
	LogFolder string
}

func InitConfig() {
	Env = &AppConfig{
		Env:               os.Getenv(keyEnv),
		Host:              os.Getenv(keyHost),
		Port:              os.Getenv(keyPort),
		AccessAllowOrigin: os.Getenv(keyAccessAllowOrigin),
		DatabaseDSN:       os.Getenv(keyDatabaseDSN),
		DataPath: &DataPathConfig{
			LogFolder: os.Getenv(keyLogFolderPath),
		},
	}
}

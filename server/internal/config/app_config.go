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
	keyDatabasePath      = "SQLALCHEMY_DATABASE_URI"
	keyLogFolderPath     = "LOG_FOLDER_PATH"
)

type AppConfig struct {
	Env               string
	Host              string
	Port              string
	AccessAllowOrigin string
	DataPath          *DataPathConfig
}

type DataPathConfig struct {
	Database  string
	LogFolder string
}

func InitConfig() {
	Env = &AppConfig{
		Env:               os.Getenv(keyEnv),
		Host:              os.Getenv(keyHost),
		Port:              os.Getenv(keyPort),
		AccessAllowOrigin: os.Getenv(keyAccessAllowOrigin),
		DataPath: &DataPathConfig{
			Database:  os.Getenv(keyDatabasePath),
			LogFolder: os.Getenv(keyLogFolderPath),
		},
	}
}

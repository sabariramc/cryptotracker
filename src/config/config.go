package config

import (
	"github.com/sabariramc/goserverbase/config"
	"github.com/sabariramc/goserverbase/utils"
)

type MasterConfig struct {
	Logger *config.LoggerConfig
	App    *config.ServerConfig
	Mongo  *config.MongoConfig
}

func NewConfig() *MasterConfig {
	return &MasterConfig{
		Logger: &config.LoggerConfig{
			Version:           utils.GetEnv("LOG_VERSION", "1.1"),
			Host:              utils.GetHostName(),
			ServiceName:       utils.GetEnv("SERVICE_NAME", "API"),
			LogLevel:          utils.GetEnvInt("LOG_LEVEL", 6),
			AuthHeaderKeyList: utils.GetEnvAsSlice("AUTH_HEADER_LIST", []string{}, ";"),
		},
		App: &config.ServerConfig{
			Host:        "0.0.0.0",
			Port:        "3000",
			ServiceName: utils.GetEnv("SERVICE_NAME", "API"),
		},
		Mongo: &config.MongoConfig{
			ConnectionString: utils.GetEnv("MONGO_URL", ""),
			DatabaseName:     utils.GetEnvMust("MONGO_DATABASENAME"),
		},
	}
}

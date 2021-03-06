package config

import (
	"fmt"

	"github.com/sabariramc/goserverbase/config"
	"github.com/sabariramc/goserverbase/utils"
	"github.com/shopspring/decimal"
)

type EmailServerConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type PriceTrackerConfig struct {
	URL string
}

type MoniterConfig struct {
	PeriodInSeconds int
}

type NotifierConfig struct {
	High         decimal.Decimal
	Low          decimal.Decimal
	EmailAddress string
}

type MasterConfig struct {
	Logger       *config.LoggerConfig
	App          *config.ServerConfig
	Mongo        *config.MongoConfig
	Email        *EmailServerConfig
	PriceTracker *PriceTrackerConfig
	Moniter      *MoniterConfig
	Notifier     *NotifierConfig
}

func NewConfig() *MasterConfig {
	high, err := decimal.NewFromString(utils.GetEnvMust("NOTIFIER_HIGH_PRICE"))
	if err != nil {
		panic(fmt.Errorf("NewConfig.InvalidNotifierHighPrice : %w", err))
	}
	low, err := decimal.NewFromString(utils.GetEnvMust("NOTIFIER_LOW_PRICE"))
	if err != nil {
		panic(fmt.Errorf("NewConfig.InvalidNotifierLowPrice : %w", err))
	}
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
			ConnectionString: utils.GetEnvMust("MONGO_URL"),
			DatabaseName:     utils.GetEnvMust("MONGO_DATABASENAME"),
		},
		Email: &EmailServerConfig{
			Host:     utils.GetEnvMust("EMAIL_HOST"),
			Port:     utils.GetEnvMust("EMAIL_PORT"),
			Username: utils.GetEnvMust("EMAIL_USERNAME"),
			Password: utils.GetEnvMust("EMAIL_PASSWORD"),
		},
		PriceTracker: &PriceTrackerConfig{
			URL: utils.GetEnvMust("PRICE_TACKER_URL"),
		},
		Moniter: &MoniterConfig{
			PeriodInSeconds: utils.GetEnvInt("MONITER_PERIOD_IN_SECONDS", 30),
		},
		Notifier: &NotifierConfig{
			High:         high,
			Low:          low,
			EmailAddress: utils.GetEnvMust("NOTIFIER_USER_EMAIL"),
		},
	}
}

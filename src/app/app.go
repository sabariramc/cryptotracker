package app

import (
	"context"

	"thinklink/src/config"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/sabariramc/goserverbase/baseapp"
	"github.com/sabariramc/goserverbase/db/mongo"
	"github.com/sabariramc/goserverbase/log"
	"github.com/sabariramc/goserverbase/log/logwriter"
	"gopkg.in/validator.v2"
)

type BitCoinTacker struct {
	*baseapp.BaseApp
	db        *mongo.Mongo
	log       *log.Logger
	validator *validator.Validator
}

func GetDefaultApp() (*BitCoinTacker, error) {
	c := config.NewConfig()
	hostParams := &log.HostParams{
		Host:        c.App.Host,
		Version:     "1.0",
		ServiceName: c.App.ServiceName,
	}
	consoleLogger := logwriter.NewConsoleWriter(*hostParams)
	lmux := log.NewSequenctialLogMultipluxer(consoleLogger)
	return GetApp(c, lmux, consoleLogger, session.Must(session.NewSession()))
}

func GetApp(c *config.MasterConfig, lMux log.LogMultipluxer, auditLog log.AuditLogWriter, awsSession *session.Session) (*BitCoinTacker, error) {
	r := &BitCoinTacker{
		BaseApp: baseapp.NewBaseApp(baseapp.ServerConfig{
			LoggerConfig: c.Logger,
			AppConfig:    c.App,
		}, lMux, auditLog),
		validator: validator.NewValidator(),
	}
	ctx := r.GetCorrelationContext(context.Background(), log.GetDefaultCorrelationParams(c.App.ServiceName))
	r.log = r.GetLogger()
	conn, err := mongo.NewMongo(ctx, r.log, *c.Mongo)
	if err != nil {
		return nil, err
	}
	r.db = conn
	r.log.Info(ctx, "App Created", nil)
	r.RegisterRoutes(ctx, r.Routes())
	r.log.Info(ctx, "Routes Registered", nil)
	r.log.Info(ctx, "Starting server on port - "+r.GetPort(), nil)
	return r, nil
}

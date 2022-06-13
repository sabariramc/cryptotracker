package app

import (
	"context"
	"sync"

	"cryptotracker/src/config"
	constant "cryptotracker/src/constants"

	"github.com/sabariramc/goserverbase/baseapp"
	"github.com/sabariramc/goserverbase/db/mongo"
	"github.com/sabariramc/goserverbase/log"
	"github.com/sabariramc/goserverbase/log/logwriter"
)

type CryptoTacker struct {
	*baseapp.BaseApp
	db                  *mongo.Mongo
	log                 *log.Logger
	priceTrackerClient  PriceTracker
	emailNotifierClient EmailNotifier
	c                   *config.MasterConfig
}

func GetDefaultApp() (*CryptoTacker, error) {
	c := config.NewConfig()
	hostParams := &log.HostParams{
		Host:        c.App.Host,
		Version:     "1.0",
		ServiceName: c.App.ServiceName,
	}
	consoleLogger := logwriter.NewConsoleWriter(*hostParams)
	lmux := log.NewSequenctialLogMultipluxer(consoleLogger)
	return GetApp(c, lmux, consoleLogger, nil, nil)
}

func GetApp(c *config.MasterConfig, lMux log.LogMultipluxer, auditLog log.AuditLogWriter, priceTrackerClient PriceTracker, emailNotifierClient EmailNotifier) (*CryptoTacker, error) {
	r := &CryptoTacker{
		BaseApp: baseapp.NewBaseApp(baseapp.ServerConfig{
			LoggerConfig: c.Logger,
			AppConfig:    c.App,
		}, lMux, auditLog),
		c:                   c,
		priceTrackerClient:  priceTrackerClient,
		emailNotifierClient: emailNotifierClient,
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

func (ct *CryptoTacker) StartJob(ctx context.Context) {
	var wg sync.WaitGroup
	defer wg.Wait()
	ch := make(chan *Price, 10)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ct.Notifier(ctx, ch)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		ct.Moniter(ctx, constant.BITCOIN, ch)
	}()
}

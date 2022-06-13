package coingecko

import (
	"context"
	"cryptotracker/src/config"
	"cryptotracker/src/model"
	"fmt"
	"net/http"
	"time"
)

type Coingecko struct {
	c config.PriceTrackerConfig
}

func (cg *Coingecko) GetCurrentPrice(ctx context.Context, coin, currency string) (*model.CryptoPrice, error) {
	return nil, nil
}
func (cg *Coingecko) GetPrice(ctx context.Context, coin, currency string, fromDate, toDate time.Time) ([]*model.CryptoPrice, error) {
	resp, err := http.Get(fmt.Sprintf("%v/coins/%v/market_chart/range?vs_currency=%v&from=%v&to=%v'", cg.c.URL, coin, currency, fromDate.Unix(), toDate.Unix()))
	if err != nil {
		return nil, fmt.Errorf("Coingecko.GetPrice : %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("CounGecko.GetPrice : %w", fmt.Errorf("Status not 200 %v", resp.StatusCode))
	}
	return nil, nil
}

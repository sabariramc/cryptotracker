package app

import (
	"context"
	constant "cryptotracker/src/constants"
	"fmt"
	"time"
)

func (ct *CryptoTacker) Moniter(ctx context.Context, symbol string, ch chan *Price) {
	defer close(ch)
	tctx, _ := context.WithTimeout(ctx, time.Second*time.Duration(ct.c.Moniter.PeriodInSeconds))
	for {
		select {
		case <-tctx.Done():
			tctx, _ = context.WithTimeout(ctx, time.Second*time.Duration(ct.c.Moniter.PeriodInSeconds))
			data, err := ct.priceTrackerClient.GetCurrentPrice(ctx, symbol, constant.USD)
			ct.log.Debug(ctx, "Price", data)
			if err != nil {
				ct.log.Error(ctx, "Error fetching price", fmt.Errorf("CryptoTacker.Moniter : %w", err))
			} else {
				ch <- data
			}
		case <-ctx.Done():
			return
		}
	}
}

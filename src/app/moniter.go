package app

import (
	"context"
	"time"
)

func (bt *BitCoinTacker) Moniter(ctx context.Context, symbol string, ch chan *Price) {
	defer close(ch)
	tctx, cancel := context.WithTimeout(ctx, time.Second*30)

	for {
		select {}
	}
}

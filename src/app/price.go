package app

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

type Price struct {
	Timestamp time.Time
	Price     decimal.Decimal
	Symbol    string
	Currency  string
}

type PriceTracker interface {
	GetCurrentPrice(ctx context.Context, symbol string, currency string) (*Price, error)
	GetPrice(ctx context.Context, symbol string, currency string, fromDate time.Time, toDate time.Time) ([]*Price, error)
}

package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type PriceResponseDTO struct {
	Timestamp time.Time       `json:"timestamp"`
	Price     decimal.Decimal `json:"price"`
	Coin      string          `json:"coin"`
}

type PriceHistoryResponseDTO struct {
	URL     string              `json:"url"`
	NextURL string              `json:"next"`
	Count   int                 `json:"count"`
	Data    []*PriceResponseDTO `json:"data"`
}

type PriceRequestDTO struct {
	Date   time.Time `json:"date"  schema:"date"`
	Limit  int       `json:"limit" schema:"limit"`
	Offset int       `json:"offset"  schema:"offset"`
}

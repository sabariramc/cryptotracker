package app

import (
	"context"
	"cryptotracker/src/model"
)

type EmailMessage struct {
	Message  string
	Coin     string
	Currency string
	Price    model.CryptoPrice
}

type EmailNotifier interface {
	Send(ctx context.Context, to string, message *EmailMessage) error
}

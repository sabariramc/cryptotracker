package app

import (
	"context"
)

type EmailMessage struct {
	CurrentPrice *Price
	Message      string
}

type EmailNotifier interface {
	Send(ctx context.Context, to string, currency string) error
}

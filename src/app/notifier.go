package app

import (
	"context"
	"cryptotracker/src/model"
	"fmt"
)

func (ct *CryptoTacker) Notifier(ctx context.Context, ch chan *model.CryptoPrice) {
	for price := range ch {
		var msg *EmailMessage
		if price.Price.Cmp(ct.c.Notifier.High) >= 0 {
			msg = &EmailMessage{
				Message: fmt.Sprintf("Price gaining above USD %v", ct.c.Notifier.High),
			}
		} else if price.Price.Cmp(ct.c.Notifier.Low) <= 0 {
			msg = &EmailMessage{
				Message: fmt.Sprintf("Price lossing below USD %v", ct.c.Notifier.High),
			}
		}
		if msg != nil {
			ct.log.Info(ctx, "Publising price alert", msg)
			msg.Price = *price
			err := ct.emailNotifierClient.Send(ctx, ct.c.Notifier.EmailAddress, msg)
			if err != nil {
				ct.log.Error(ctx, "Error sending mail", fmt.Errorf("CryptoTacker.Notifier : %w", err))
			}
		}
	}
}

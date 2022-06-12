package model

import (
	"context"
	"fmt"
	"time"

	"github.com/sabariramc/goserverbase/constant"
	"github.com/sabariramc/goserverbase/db/mongo"
	"github.com/shopspring/decimal"
)

type CryptoPrice struct {
	Timestamp time.Time       `json:"timestamp" bson:"timestamp"`
	Price     decimal.Decimal `json:"price" bson:"price"`
}

type CryptoTracker struct {
	mongo.BaseMongoModel `bson:",inline"`
	CoinSymbol           string `json:"coinSymbol" bson:"coinSymbol"`
	Date                 time.Time
	PriceHistory         []CryptoPrice `json:"priceList" bson:"PriceHistory"`
}

func (a *CryptoTracker) Create(ctx context.Context, db *mongo.Mongo) string {
	coll := db.NewCollection("CryptoTracker")
	actor := ctx.Value(constant.ActorIdKey).(string)
	a.SetCreateParam(actor)
	_, err := coll.InsertOne(ctx, a)
	if err != nil {
		return fmt.Errorf("CryptoTracker.Create: %w", err)
	}
	return nil
}

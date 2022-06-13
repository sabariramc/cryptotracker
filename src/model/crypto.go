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
	Coin                 string         `json:"coin" bson:"coin"`
	Date                 time.Time      `json:"date" bson:"date"`
	Currency             string         `json:"currency" bson:"currency"`
	PriceHistory         []*CryptoPrice `json:"priceHistory" bson:"priceHistory"`
}

func (ct *CryptoTracker) Create(ctx context.Context, db *mongo.Mongo) error {
	coll := db.NewCollection("CryptoTracker")
	actor := ctx.Value(constant.ActorIdKey).(string)
	ct.SetCreateParam(actor)
	_, err := coll.InsertOne(ctx, ct)
	if err != nil {
		return fmt.Errorf("CryptoTracker.Create: %w", err)
	}
	return nil
}

func (ct *CryptoTracker) Get(ctx context.Context, db *mongo.Mongo, coin string, onDate time.Time) error {
	coll := db.NewCollection("CryptoTracker")
	cur := coll.FindOne(ctx, map[string]interface{}{"coin": coin, "date": onDate})
	return cur.Decode(ct)
}

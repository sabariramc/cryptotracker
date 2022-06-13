package app

import (
	"context"
	"cryptotracker/src/constant"
	"cryptotracker/src/model"
	"cryptotracker/src/model/dto"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/schema"
	"github.com/sabariramc/goserverbase/baseapp"
	"github.com/sabariramc/goserverbase/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

var timeConverter = func(value string) reflect.Value {
	if v, err := time.Parse("22-12-2006", value); err == nil {
		return reflect.ValueOf(v)
	}
	return reflect.Value{}
}

type PriceTracker interface {
	GetCurrentPrice(ctx context.Context, coin, currency string) (*model.CryptoPrice, error)
	GetPrice(ctx context.Context, coin, currency string, fromDate, toDate time.Time) ([]*model.CryptoPrice, error)
}

func (ct *CryptoTacker) GetPrice() http.HandlerFunc {
	return ct.JSONResponder(nil, func(r *http.Request) (statusCode int, res interface{}, err error) {
		var qp dto.PriceRequestDTO
		decoder := schema.NewDecoder()
		decoder.RegisterConverter(time.Time{}, timeConverter)
		err = decoder.Decode(&qp, r.URL.Query())
		if err != nil {
			return http.StatusBadRequest, nil, fmt.Errorf("CryptoTacker.GetPrice : %w", err)
		}
		ctx := r.Context()
		params := baseapp.GetPathParams(ctx, ct.log, r)
		coinParam := params[0].Value
		coin, ok := constant.SUPPORTED_MAP[coinParam]
		if !ok {
			return http.StatusNotFound, nil, fmt.Errorf("Invalid path")
		}
		var v model.CryptoTracker
		day, month, year := time.Now().Date()
		if qp.Date.After(time.Date(day, month, year, 0, 0, 0, 0, time.Local)) {
			return http.StatusOK, nil, nil
		} else {
			err = v.Get(ctx, ct.db, coin, qp.Date)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					data, err := ct.priceTrackerClient.GetPrice(ctx, coin, constant.USD, qp.Date, qp.Date.Add(time.Hour*24))
					if err != nil {
						return http.StatusBadGateway, nil, fmt.Errorf("CryptoTacker.GetPrice : %w", err)
					}
					v.PriceHistory = data
					v.Coin = coin
					v.Currency = constant.USD
					err = v.Create(ctx, ct.db)
					if err != nil {
						return http.StatusInternalServerError, nil, fmt.Errorf("CryptoTacker.GetPrice : %w", err)
					}
				} else {
					return http.StatusInternalServerError, nil, fmt.Errorf("CryptoTacker.GetPrice : %w", err)
				}
			}
			res := &dto.PriceHistoryResponseDTO{}
			n := len(v.PriceHistory)
			res.Count = n
			res.URL = fmt.Sprintf("<http://%v/api/prices/btc?date=%v&offset=%v&limit=%v>", ct.GetPort(), qp.Date.Format("21-12-2021"), qp.Offset, qp.Limit)
			if qp.Offset > n {
				res.NextURL = ""
				res.Data = make([]*dto.PriceResponseDTO, 0)
			} else {
				res.NextURL = fmt.Sprintf("<http://%v/api/prices/btc?date=%v&offset=%v&limit=%v>", ct.GetPort(), qp.Date.Format("21-12-2021"), qp.Offset+qp.Limit, qp.Limit)
				resCount := n - qp.Offset
				if resCount > qp.Limit {
					resCount = qp.Limit
				}
				res.Data = make([]*dto.PriceResponseDTO, resCount)
				for i, j := qp.Offset, 0; i < qp.Offset+resCount; i, j = i+1, j+1 {
					resPart := &dto.PriceResponseDTO{}
					utils.JsonTransformer(v.PriceHistory[i], resPart)
					res.Data[j] = resPart
				}
			}

			return http.StatusOK, res, nil
		}
	})
}

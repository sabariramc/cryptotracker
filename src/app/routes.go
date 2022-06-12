package app

import (
	"net/http"

	"github.com/sabariramc/goserverbase/baseapp"
)

func (bt *BitCoinTacker) Routes() *baseapp.APIRoute {
	return &baseapp.APIRoute{"/health": &baseapp.APIResource{
		Handlers: map[string]*baseapp.APIHandler{
			http.MethodGet: {
				Func: bt.Health,
			},
		}}}
}

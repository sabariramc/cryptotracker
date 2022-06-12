package app

import "net/http"

func (bt *BitCoinTacker) Health(w http.ResponseWriter, r *http.Request) {
	bt.log.Debug(r.Context(), "Startup Health Check", nil)
	w.WriteHeader(http.StatusNoContent)
}

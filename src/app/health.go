package app

import "net/http"

func (ct *CryptoTacker) Health(w http.ResponseWriter, r *http.Request) {
	ct.log.Debug(r.Context(), "Startup Health Check", nil)
	w.WriteHeader(http.StatusNoContent)
}

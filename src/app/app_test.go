package app_test

import (
	"cryptotracker/src/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sabariramc/goserverbase/utils"
	"gotest.tools/assert"
)

func TestApp(t *testing.T) {
	s, err := app.GetDefaultApp()
	assert.NilError(t, err)
	s.GetLogger()
}

func TestHealth(t *testing.T) {
	s, err := app.GetDefaultApp()
	assert.NilError(t, err)
	req := httptest.NewRequest("GET", "/health", nil)
	req.Header.Set("x-api-key", utils.GetEnv("TEST_API_KEY", ""))
	w := httptest.NewRecorder()
	s.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Result().StatusCode)
}

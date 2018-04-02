// +build integration
package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aerogear/aerogear-app-metrics/internal/setup"
	"github.com/aerogear/aerogear-app-metrics/pkg/config"
	"github.com/stretchr/testify/assert"
)

func setupTestServer() *httptest.Server {
	config := config.GetConfig()
	metricsDao := setup.InitDao(config)
	router := setup.InitRouter(metricsDao)

	return httptest.NewServer(router)
}

func TestPingEndpoint(t *testing.T) {
	s := setupTestServer()
	defer s.Close()

	res, err := http.Get(s.URL + "/ping")
	assert.NoError(t, err, "did not expect an error on GET /ping")
	assert.Equal(t, 200, res.StatusCode)
}

func TestHealtzEndpoint(t *testing.T) {
	s := setupTestServer()
	defer s.Close()

	res, err := http.Get(s.URL + "/healthz")
	assert.NoError(t, err, "did not expect an error on GET /healthz")
	assert.Equal(t, 200, res.StatusCode)
}

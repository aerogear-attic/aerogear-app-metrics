// +build integration
package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aerogear/aerogear-app-metrics/internal/setup"
	"github.com/aerogear/aerogear-app-metrics/pkg/config"
	"github.com/aerogear/aerogear-app-metrics/pkg/mobile"
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

func TestMetricsEndpoint(t *testing.T) {
	s := setupTestServer()
	defer s.Close()

	cases := []struct {
		Name     string
		Expected int
		Payload  mobile.Metric
	}{
		{
			Name:     "Valid Init Metric should be registered",
			Expected: 204,
			Payload:  GetValidInitMetric(),
		},
		{
			Name:     "Init Metric without 'app' should error out",
			Expected: 400,
			Payload:  GetNoAppInitMetric(),
		},
		{
			Name:     "Init Metric without 'device' should error out",
			Expected: 400,
			Payload:  GetNoDeviceInitMetric(),
		},
		{
			Name:     "Valid Security Metric should be registered",
			Expected: 204,
			Payload:  GetValidSecurityMetric(),
		},
		{
			Name:     "Incomplete Security Metric should error out",
			Expected: 400,
			Payload:  GetIncompleteSecurityMetric(),
		},
		{
			Name:     "Empty Security Metric should error out",
			Expected: 400,
			Payload:  GetEmptySecurityMetric(),
		},
		{
			Name:     "Large clientIds should error out",
			Expected: 400,
			Payload:  GetLargeClientIdMetric(),
		},
	}
	for _, c := range cases {
		var buffer *bytes.Buffer
		buffer = new(bytes.Buffer)
		err := json.NewEncoder(buffer).Encode(c.Payload)
		assert.NoError(t, err, "did not expect an error marshaling metric")

		res, err := http.Post(s.URL+"/metrics", "application/json", buffer)
		assert.NoError(t, err, "did not expect an error posting to /metrics")
		defer res.Body.Close()

		assert.Equalf(t, c.Expected, res.StatusCode, "Failure in test %v\nexpected: %v\ngot: %v", c.Name, c.Expected, res.StatusCode)
	}
}

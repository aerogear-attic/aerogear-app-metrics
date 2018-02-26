package web

import (
	"testing"

	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/aerogear/aerogear-app-metrics/pkg/mobile"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

type MockMetricsService struct {
	mock.Mock
}

func (m *MockMetricsService) Create(metric mobile.Metric) (mobile.Metric, error) {
	args := m.Called(metric)
	return args.Get(0).(mobile.Metric), args.Error(1)
}

func setupMetricsHandler(service MetricsServiceInterface) *httptest.Server {
	router := NewRouter()
	metricsHandler := NewMetricsHandler(service)
	MetricsRoute(router, metricsHandler)
	return httptest.NewServer(router)
}

func TestMetricsEndpointShouldPassReceivedDataToMetricsService(t *testing.T) {
	metric := mobile.Metric{
		ClientTimestamp: 1234,
		ClientId:        "client123",
		Data: &mobile.MetricData{
			App: &mobile.AppMetric{
				ID:         "deadbeef",
				SDKVersion: "1.2.3",
				AppVersion: "27",
			},
			Device: &mobile.DeviceMetric{
				Platform:        "android",
				PlatformVersion: "19",
			},
		},
	}

	byteBuffer := new(bytes.Buffer)
	err := json.NewEncoder(byteBuffer).Encode(metric)
	assert.Nil(t, err, "did not expect an error marshaling metric")

	mockMetricsService := new(MockMetricsService)
	mockMetricsService.On("Create", metric).Return(metric, nil)

	s := setupMetricsHandler(mockMetricsService)
	defer s.Close()

	res, err := http.Post(s.URL+"/metrics", "application/json", byteBuffer)
	assert.Nil(t, err, "did not expect an error posting metrics")

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)

	assert.Equal(t, 200, res.StatusCode)

	mockMetricsService.AssertExpectations(t)
}

func TestMetricsEndpointShouldReturn500WhenThereIsAnErrorInMetricsService(t *testing.T) {
	metric := mobile.Metric{
		ClientTimestamp: 1234,
		ClientId:        "client123",
		Data: &mobile.MetricData{
			App: &mobile.AppMetric{
				ID:         "deadbeef",
				SDKVersion: "1.2.3",
				AppVersion: "27",
			},
			Device: &mobile.DeviceMetric{
				Platform:        "android",
				PlatformVersion: "19",
			},
		},
	}

	byteBuffer := new(bytes.Buffer)
	err := json.NewEncoder(byteBuffer).Encode(metric)
	assert.Nil(t, err, "did not expect an error marshaling metric")

	mockMetricsService := new(MockMetricsService)
	mockMetricsService.On("Create", metric).Return(mobile.Metric{}, errors.New("Metrics service error!")) // mock the service so that it returns an error

	s := setupMetricsHandler(mockMetricsService)
	defer s.Close()

	res, err := http.Post(s.URL+"/metrics", "application/json", byteBuffer)
	assert.Nil(t, err, "did not expect an error posting metrics")

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)

	assert.Equal(t, 500, res.StatusCode)

	mockMetricsService.AssertExpectations(t)
}

func TestMetricsEndpointShouldNotInteractWithMetricsServiceWhenRequestBodyIsEmpty(t *testing.T) {
	mockMetricsService := new(MockMetricsService)

	s := setupMetricsHandler(mockMetricsService)
	defer s.Close()

	res, err := http.Post(s.URL+"/metrics", "application/json", nil) // empty request body
	assert.Nil(t, err, "did not expect an error posting metrics")

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)

	assert.Equal(t, 400, res.StatusCode)

	mockMetricsService.AssertNotCalled(t, "Create")
}

func TestMetricsEndpointShouldNotInteractWithMetricsServiceWhenRequestBodyIsInvalidJSON(t *testing.T) {
	mockMetricsService := new(MockMetricsService)

	s := setupMetricsHandler(mockMetricsService)
	defer s.Close()

	byteBuffer := new(bytes.Buffer)
	byteBuffer.WriteString("nonsense") // invalid JSON

	res, err := http.Post(s.URL+"/metrics", "application/json", byteBuffer)
	assert.Nil(t, err, "did not expect an error posting metrics")

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)

	assert.Equal(t, 400, res.StatusCode)

	mockMetricsService.AssertNotCalled(t, "Create")
}

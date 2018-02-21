package web

import (
	"testing"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"github.com/aerogear/aerogear-metrics-api/pkg/mobile"
	"bytes"
	"github.com/stretchr/testify/mock"
	"github.com/pkg/errors"
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
		App: mobile.AppMetric{
			ID:         "deadbeef",
			SDKVersion: "1.2.3",
			AppVersion: "27",
		},
		Device: mobile.DeviceMetric{
			Platform:        "android",
			PlatformVersion: "19",
		},
	}

	byteBuffer := new(bytes.Buffer)
	err := json.NewEncoder(byteBuffer).Encode(metric)

	if err != nil {
		t.Fatal("unable to marshal metric", err)
	}

	mockMetricsService := new(MockMetricsService)
	mockMetricsService.On("Create", metric).Return(metric, nil)

	s := setupMetricsHandler(mockMetricsService)
	defer s.Close()

	res, err := http.Post(s.URL+"/metrics", "application/json", byteBuffer)
	if err != nil {
		t.Fatal("did not expect an error posting metrics", err)
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		t.Fatal("expected a 200 status")
	}

	mockMetricsService.AssertExpectations(t)
}

func TestMetricsEndpointShouldReturn500WhenThereIsAnErrorInMetricsService(t *testing.T) {
	metric := mobile.Metric{
		ClientTimestamp: 1234,
		ClientId:        "client123",
		App: mobile.AppMetric{
			ID:         "deadbeef",
			SDKVersion: "1.2.3",
			AppVersion: "27",
		},
		Device: mobile.DeviceMetric{
			Platform:        "android",
			PlatformVersion: "19",
		},
	}

	byteBuffer := new(bytes.Buffer)
	err := json.NewEncoder(byteBuffer).Encode(metric)

	if err != nil {
		t.Fatal("unable to marshal metric", err)
	}

	mockMetricsService := new(MockMetricsService)
	mockMetricsService.On("Create", metric).Return(mobile.Metric{}, errors.New("Metrics service error!"))		// mock the service so that it returns an error

	s := setupMetricsHandler(mockMetricsService)
	defer s.Close()

	res, err := http.Post(s.URL+"/metrics", "application/json", byteBuffer)
	if err != nil {
		t.Fatal("did not expect an error posting metrics", err)
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)

	if res.StatusCode != 500 {
		t.Fatal("expected a 500 status")
	}

	mockMetricsService.AssertExpectations(t)
}

func TestMetricsEndpointShouldNotInteractWithMetricsServiceWhenRequestBodyIsEmpty(t *testing.T) {
	mockMetricsService := new(MockMetricsService)

	s := setupMetricsHandler(mockMetricsService)
	defer s.Close()

	res, err := http.Post(s.URL+"/metrics", "application/json", nil) // empty request body
	if err != nil {
		t.Fatal("did not expect an error posting metrics", err)
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)

	if res.StatusCode != 400 {
		t.Fatal("expected a 400 status")
	}

	mockMetricsService.AssertNotCalled(t, "Create")
}

func TestMetricsEndpointShouldNotInteractWithMetricsServiceWhenRequestBodyIsInvalidJSON(t *testing.T) {
	mockMetricsService := new(MockMetricsService)

	s := setupMetricsHandler(mockMetricsService)
	defer s.Close()

	byteBuffer := new(bytes.Buffer)
	byteBuffer.WriteString("nonsense") // invalid JSON

	res, err := http.Post(s.URL+"/metrics", "application/json", byteBuffer)
	if err != nil {
		t.Fatal("did not expect an error posting metrics", err)
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)

	if res.StatusCode != 400 {
		t.Fatal("expected a 400 status")
	}

	mockMetricsService.AssertNotCalled(t, "Create")
}

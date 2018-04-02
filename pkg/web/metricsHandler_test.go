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
	"github.com/aerogear/aerogear-app-metrics/pkg/test"
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

func postMetric(t *testing.T, s *httptest.Server, metric *mobile.Metric) (*http.Response, error) {
	var buffer *bytes.Buffer
	buffer = new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(metric)
	assert.Nil(t, err, "did not expect an error marshaling metric")

	return postBuffer(t, s, buffer)
}

func postBuffer(t *testing.T, s *httptest.Server, buffer *bytes.Buffer) (res *http.Response, err error) {
	if buffer != nil {
		res, err = http.Post(s.URL+"/metrics", "application/json", buffer)
	} else {
		res, err = http.Post(s.URL+"/metrics", "application/json", nil)
	}
	assert.Nil(t, err, "did not expect an error posting metrics")
	_, err = ioutil.ReadAll(res.Body)
	assert.Nil(t, err, "did not expect an error reading the response body")
	return res, err
}

func TestMetricsEndpointShouldPassReceivedDataToMetricsService(t *testing.T) {
	metric := test.GetValidInitMetric()

	mockMetricsService := new(MockMetricsService)
	mockMetricsService.On("Create", metric).Return(metric, nil)
	s := setupMetricsHandler(mockMetricsService)
	defer s.Close()

	res, _ := postMetric(t, s, &metric)
	defer res.Body.Close()
	assert.Equal(t, 204, res.StatusCode)

	mockMetricsService.AssertExpectations(t)
}

func TestMetricsEndpointShouldReturn500WhenThereIsAnErrorInMetricsService(t *testing.T) {
	metric := test.GetValidInitMetric()

	mockMetricsService := new(MockMetricsService)
	mockMetricsService.On("Create", metric).Return(mobile.Metric{}, errors.New("metrics service error")) // mock the service so that it returns an error

	s := setupMetricsHandler(mockMetricsService)
	defer s.Close()

	res, _ := postMetric(t, s, &metric)
	defer res.Body.Close()
	assert.Equal(t, 500, res.StatusCode)

	mockMetricsService.AssertExpectations(t)
}

func TestMetricsEndpointShouldNotInteractWithMetricsServiceWhenRequestBodyIsEmpty(t *testing.T) {
	mockMetricsService := new(MockMetricsService)

	s := setupMetricsHandler(mockMetricsService)
	defer s.Close()

	res, _ := postBuffer(t, s, nil)
	defer res.Body.Close()

	assert.Equal(t, 400, res.StatusCode)

	mockMetricsService.AssertNotCalled(t, "Create")
}

func TestMetricsEndpointShouldNotInteractWithMetricsServiceWhenRequestBodyIsInvalidJSON(t *testing.T) {
	mockMetricsService := new(MockMetricsService)

	s := setupMetricsHandler(mockMetricsService)
	defer s.Close()

	byteBuffer := new(bytes.Buffer)
	byteBuffer.WriteString("nonsense") // invalid JSON

	res, _ := postBuffer(t, s, byteBuffer)
	defer res.Body.Close()

	assert.Equal(t, 400, res.StatusCode)

	mockMetricsService.AssertNotCalled(t, "Create")
}

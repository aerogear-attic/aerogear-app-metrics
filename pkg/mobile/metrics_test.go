package mobile

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MetricsDAOMock struct {
	mock.Mock
}

func (m *MetricsDAOMock) Create(clientId string, metricsData []byte) error {
	args := m.Called(clientId, metricsData)
	return args.Error(0)
}

func newTestMetricsService() (*MetricsDAOMock, *MetricsService) {
	mdaoMock := MetricsDAOMock{}

	ms := NewMetricsService(&mdaoMock)

	return &mdaoMock, ms
}

func TestCreateCallsDAOWithCorrectArgs(t *testing.T) {

	metric := Metric{
		ClientId: "org.aerogear.metrics.tests",
		Data: &MetricData{
			App: &AppMetric{
				ID:         "12345678",
				SDKVersion: "1.0.0",
				AppVersion: "1",
			},
			Device: &DeviceMetric{
				Platform:        "Android",
				PlatformVersion: "27",
			},
		},
	}
	expectedMetricsData, err := json.Marshal(metric.Data)

	if err != nil {
		t.Errorf("could not encode metric object to JSON")
	}

	mdaoMock, ms := newTestMetricsService()

	mdaoMock.On("Create", metric.ClientId, expectedMetricsData).Return(nil)

	res, err := ms.Create(metric)

	if err != nil {
		t.Errorf("Metrics Service should not have returned an error")
	}

	if reflect.DeepEqual(reflect.ValueOf(metric), reflect.ValueOf(res)) {
		t.Errorf("failed")
	}

	mdaoMock.AssertExpectations(t)
}

func TestCreateReturnsErrorFromDAO(t *testing.T) {

	metric := Metric{
		ClientId: "org.aerogear.metrics.tests",
		Data: &MetricData{
			App: &AppMetric{
				ID:         "12345678",
				SDKVersion: "1.0.0",
				AppVersion: "1",
			},
			Device: &DeviceMetric{
				Platform:        "Android",
				PlatformVersion: "27",
			},
		},
	}
	expectedMetricsData, err := json.Marshal(metric.Data)

	if err != nil {
		t.Errorf("could not encode metric object to JSON")
	}

	mdaoMock, ms := newTestMetricsService()

	daoError := errors.New("problem connecting to db")
	mdaoMock.On("Create", metric.ClientId, expectedMetricsData).Return(daoError)

	_, err = ms.Create(metric)

	if err.Error() != daoError.Error() {
		t.Errorf("Metrics Service did not return the error from the DAO")
	}

	mdaoMock.AssertExpectations(t)
}

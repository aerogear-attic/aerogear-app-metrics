package web

import (
	"errors"
	"testing"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockHealthCheckable implements HealthCheckable interface
type MockHealthCheckable struct {
	mock.Mock
}

// IsHealthy provides a mock function with given fields:
func (_m *MockHealthCheckable) IsHealthy() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func TestHealthz(t *testing.T) {
	checkable := &MockHealthCheckable{}

	healthHandler := NewHealthHandler(checkable)

	assert.HTTPSuccess(t, healthHandler.Healthz, "GET", "/healthz", nil, nil)
	checkable.AssertExpectations(t)
}

func TestHealthzPing(t *testing.T) {
	checkable := &MockHealthCheckable{}

	checkable.On("IsHealthy").Return(true, nil).Times(1)
	checkable.On("IsHealthy").Return(false, errors.New("db offline")).Times(1)

	healthHandler := NewHealthHandler(checkable)

	assert.HTTPSuccess(t, healthHandler.Ping, "GET", "/healthz/ping", nil, nil)
	assert.HTTPError(t, healthHandler.Ping, "GET", "/healthz/ping", nil, nil)

	checkable.AssertExpectations(t)
}

func expectOkayHealthResponse(handler http.HandlerFunc, url string, t *testing.T) {
	request := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	handler(w, request)

	res := w.Result()
	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	response := &healthResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Fatal("failed to unmarshal response body", err)
	}

	assert.Equal(t, 200, res.StatusCode)
}

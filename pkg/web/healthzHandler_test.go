package web

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockHealthCheckable implements HealthCheckable interface
type MockHealthCheckable struct {
	mock.Mock
}

// IsHealthy provides a mock function with given fields:
func (_m *MockHealthCheckable) IsHealthy() error {
	args := _m.Called()
	return args.Error(0)
}

func TestPing(t *testing.T) {
	checkable := &MockHealthCheckable{}
	healthHandler := NewHealthHandler(checkable)

	assert.HTTPSuccess(t, healthHandler.Ping, "GET", "/ping", nil, nil)
}

func TestHealthzPing(t *testing.T) {
	checkable := &MockHealthCheckable{}

	healthHandler := NewHealthHandler(checkable)

	checkable.On("IsHealthy").Return(nil).Once()
	assert.HTTPSuccess(t, healthHandler.Healthz, "GET", "/healthz", nil, nil)

	checkable.On("IsHealthy").Return(errors.New("db offline")).Once()
	assert.HTTPError(t, healthHandler.Healthz, "GET", "/healthz", nil, nil)

	checkable.AssertExpectations(t)
}

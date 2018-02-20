package web

import (
	"errors"
	"testing"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/aerogear/aerogear-metrics-api/pkg/mock_web"
	"github.com/golang/mock/gomock"
)

func TestHealthz(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPingable := mock_web.NewMockPingable(mockCtrl)

	router := NewRouter()
	healthHandler := NewHealthHandler(mockPingable)
	HealthzRoute(router, healthHandler)

	expectOkayHealthResponse(router, "/healthz", t)
}

func TestHealthzPing(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPingable := mock_web.NewMockPingable(mockCtrl)

	router := NewRouter()
	healthHandler := NewHealthHandler(mockPingable)

	HealthzRoute(router, healthHandler)

	gomock.InOrder(
		mockPingable.EXPECT().Ping().Return(nil),
		mockPingable.EXPECT().Ping().Return(errors.New("no db")),
	)

	expectOkayHealthResponse(router, "/healthz/ping", t)
	expectErrorResponse(router, "/healthz/ping", t)
}

func expectOkayHealthResponse(router http.Handler, url string, t *testing.T) {
	s := httptest.NewServer(router)
	defer s.Close()

	res, err := http.Get(s.URL + url)
	if err != nil {
		t.Fatal("did not expect an error", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	response := &healthResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Fatal("failed to unmarshal response body", err)
	}

	if response.Status != "ok" {
		t.Fatal("expected an ok status")
	}
}

func expectErrorResponse(router http.Handler, url string, t *testing.T) {
	s := httptest.NewServer(router)
	defer s.Close()

	res, err := http.Get(s.URL + url)
	if err != nil {
		t.Fatal("did not expect an error", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 400 {
		t.Fatal("expected an error HTTP status", res.StatusCode)
	}
}

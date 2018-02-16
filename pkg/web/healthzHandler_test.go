package web

import (
	"testing"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func TestHealthz(t *testing.T) {
	router := NewRouter()
	healthHandler := NewHealthHandler()
	HealthzRoute(router, healthHandler)
	s := httptest.NewServer(router)
	defer s.Close()
	res, err := http.Get(s.URL + "/healthz")
	if err != nil {
		t.Fatal("did not expect an error getting healthz", err)
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

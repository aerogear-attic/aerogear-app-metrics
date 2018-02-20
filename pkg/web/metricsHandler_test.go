package web

import (
	"testing"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"github.com/aerogear/aerogear-metrics-api/pkg/mobile"
	"bytes"
	"reflect"
)

type StubMetricsService struct {
	recorded            mobile.Metric
	createMethodInvoked bool
}

func (sms *StubMetricsService) Create(metric mobile.Metric) (mobile.Metric, error) {
	sms.recorded = metric
	sms.createMethodInvoked = true
	return metric, nil
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

	stubMetricsService := &StubMetricsService{}

	router := NewRouter()
	metricsHandler := NewMetricsHandler(stubMetricsService)
	MetricsRoute(router, metricsHandler)
	s := httptest.NewServer(router)
	defer s.Close()
	res, err := http.Post(s.URL+"/metrics", "application/json", byteBuffer)
	if err != nil {
		t.Fatal("did not expect an error posting metrics", err)
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		t.Fatal("expected an 200 status")
	}

	if !stubMetricsService.createMethodInvoked {
		t.Fatal("nothing sent to metrics service")
	}

	if !reflect.DeepEqual(stubMetricsService.recorded, metric) {
		t.Fatal("unexpected metrics sent to metrics service")
	}

}

func TestMetricsEndpointShouldNotInteractWithMetricsServiceWhenRequestBodyIsEmpty(t *testing.T) {
	stubMetricsService := &StubMetricsService{}

	router := NewRouter()
	metricsHandler := NewMetricsHandler(stubMetricsService)
	MetricsRoute(router, metricsHandler)
	s := httptest.NewServer(router)
	defer s.Close()
	res, err := http.Post(s.URL+"/metrics", "application/json", nil)
	if err != nil {
		t.Fatal("did not expect an error posting metrics", err)
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)

	if res.StatusCode != 400 {
		t.Fatal("expected an 400 status")
	}

	expected := mobile.Metric{}

	if stubMetricsService.createMethodInvoked {
		t.Fatal("shouldn't have interacted with metrics service")
	}

	if stubMetricsService.recorded != expected {
		t.Fatal("unexpected metrics sent to metrics service")
	}
}

func TestMetricsEndpointShouldNotInteractWithMetricsServiceWhenRequestBodyIsInvalidJSON(t *testing.T) {
	stubMetricsService := &StubMetricsService{}

	router := NewRouter()
	metricsHandler := NewMetricsHandler(stubMetricsService)
	MetricsRoute(router, metricsHandler)
	s := httptest.NewServer(router)
	defer s.Close()

	byteBuffer := new(bytes.Buffer)
	byteBuffer.WriteString("nonsense")

	res, err := http.Post(s.URL+"/metrics", "application/json", byteBuffer)
	if err != nil {
		t.Fatal("did not expect an error posting metrics", err)
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)

	if res.StatusCode != 400 {
		t.Fatal("expected an 400 status")
	}

	expected := mobile.Metric{}

	if stubMetricsService.createMethodInvoked {
		t.Fatal("shouldn't have interacted with metrics service")
	}

	if stubMetricsService.recorded != expected {
		t.Fatal("unexpected metrics sent to metrics service")
	}
}

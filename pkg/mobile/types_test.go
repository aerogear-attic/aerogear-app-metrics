package mobile

import (
	"fmt"
	"strings"
	"testing"
)

func TestMetricValidate(t *testing.T) {

	testCases := []struct {
		Name           string
		Metric         Metric
		Valid          bool
		ExpectedReason string
	}{
		{
			Name:           "Empty metric should be invalid",
			Metric:         Metric{},
			Valid:          false,
			ExpectedReason: "missing clientId in payload",
		},
		{
			Name:           "Metric with no clientId should be invalid",
			Metric:         Metric{Data: &MetricData{App: &AppMetric{SDKVersion: "1"}}},
			Valid:          false,
			ExpectedReason: "missing clientId in payload",
		},
		{
			Name:           "Metric with long clientId should be invalid",
			Metric:         Metric{ClientId: strings.Join(make([]string, clientIdMaxLength+10), "a"), Data: &MetricData{App: &AppMetric{SDKVersion: "1"}}},
			Valid:          false,
			ExpectedReason: fmt.Sprintf("clientId exceeded maximum length of %v", clientIdMaxLength),
		},
		{
			Name:           "Metric with no Data should be invalid",
			Metric:         Metric{ClientId: "org.aerogear.metrics.testing"},
			Valid:          false,
			ExpectedReason: "missing metrics data in payload",
		},
		{
			Name:           "Metric with empty Data should be invalid",
			Metric:         Metric{ClientId: "org.aerogear.metrics.testing", Data: &MetricData{}},
			Valid:          false,
			ExpectedReason: "missing metrics data in payload",
		},
		{
			Name:           "Metric with ClientId and Some Data should be valid",
			Metric:         Metric{ClientId: "org.aerogear.metrics.testing", Data: &MetricData{App: &AppMetric{SDKVersion: "1"}}},
			Valid:          true,
			ExpectedReason: "",
		},
	}

	for _, tc := range testCases {
		valid, reason := tc.Metric.Validate()

		if valid != tc.Valid {
			t.Errorf("case failed: %s. Expected: %v, got %v", tc.Name, tc.Valid, valid)
		}

		if reason != tc.ExpectedReason {
			t.Errorf("case failed: %s. Expected: %v, got %v", tc.Name, tc.ExpectedReason, reason)
		}
	}
}

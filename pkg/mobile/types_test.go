package mobile

import (
	"fmt"
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
			fmt.Errorf("case failed: %s. Expected: %v, got %v", tc.Name, tc.Valid, valid)
		}

		if reason != tc.ExpectedReason {
			fmt.Errorf("case failed: %s. Expected: %v, got %v", tc.Name, tc.ExpectedReason, reason)
		}
	}
}

package mobile_test

import (
	"fmt"
	"testing"

	"github.com/aerogear/aerogear-app-metrics/pkg/mobile"
	"github.com/aerogear/aerogear-app-metrics/pkg/test"
)

func TestMetricValidate(t *testing.T) {
	testCases := []struct {
		Name           string
		MetricBuilder  func() mobile.Metric
		Valid          bool
		ExpectedReason string
	}{
		{
			Name:           "Empty metric should be invalid",
			MetricBuilder:  test.GetEmptyMetric,
			Valid:          false,
			ExpectedReason: mobile.MissingClientIdError,
		},
		{
			Name:           "Metric with no clientId should be invalid",
			MetricBuilder:  test.GetNoClientIdMetric,
			Valid:          false,
			ExpectedReason: mobile.MissingClientIdError,
		},
		{
			Name:           "Metric with long clientId should be invalid",
			MetricBuilder:  test.GetLargeClientIdMetric,
			Valid:          false,
			ExpectedReason: mobile.ClientIdLengthError,
		},
		{
			Name:           "Metric with no Data should be invalid",
			MetricBuilder:  test.GetNoDataMetric,
			Valid:          false,
			ExpectedReason: mobile.MissingDataError,
		},
		{
			Name:           "Metric with empty Data should be invalid",
			MetricBuilder:  test.GetEmptyDataMetric,
			Valid:          false,
			ExpectedReason: mobile.MissingDataError,
		},
		{
			Name:           "Metric with ClientId and Some Data should be valid",
			MetricBuilder:  test.GetValidInitMetric,
			Valid:          true,
			ExpectedReason: "",
		},
		{
			Name:           "Metric with bad timestamp should be invalid",
			MetricBuilder:  test.GetMetricWithInvalidTimestamp,
			Valid:          false,
			ExpectedReason: mobile.InvalidTimestampError,
		},
		{
			Name:           "Metric with valid timestamp should be valid",
			MetricBuilder:  test.GetMetricWithTimestamp,
			Valid:          true,
			ExpectedReason: "",
		},
		{
			Name:           "Security Metrics with missing id field should be invalid",
			MetricBuilder:  test.GetNoIdSecurityMetric,
			Valid:          false,
			ExpectedReason: fmt.Sprintf(mobile.SecurityMetricMissingIdError, 0),
		},
		{
			Name:           "Security Metrics with missing name field should be invalid",
			MetricBuilder:  test.GetNoNameSecurityMetric,
			Valid:          false,
			ExpectedReason: fmt.Sprintf(mobile.SecurityMetricMissingNameError, 0),
		},
		{
			Name:           "Security Metrics with missing passed field should be invalid",
			MetricBuilder:  test.GetNoPassedSecurityMetric,
			Valid:          false,
			ExpectedReason: fmt.Sprintf(mobile.SecurityMetricMissingPassedError, 0),
		},
		{
			Name:           "Empty Security Metrics slice should be invalid",
			MetricBuilder:  test.GetEmptySecurityMetric,
			Valid:          false,
			ExpectedReason: mobile.SecurityMetricsEmptyError,
		},
		{
			Name:           "Security Metrics slice with length > max length should be valid",
			MetricBuilder:  test.GetOverfilledSecurityMetric,
			Valid:          false,
			ExpectedReason: mobile.SecurityMetricsLengthError,
		},
	}

	for _, tc := range testCases {
		metric := tc.MetricBuilder()
		valid, reason := metric.Validate()

		if valid != tc.Valid {
			t.Errorf("case failed: %s. Expected: %v, got %v", tc.Name, tc.Valid, valid)
		}

		if reason != tc.ExpectedReason {
			t.Errorf("case failed: %s. Expected: %v, got %v", tc.Name, tc.ExpectedReason, reason)
		}
	}
}

package mobile

import (
	"fmt"
	"strings"
	"testing"
)

func TestMetricValidate(t *testing.T) {

	securityMetricId := "org.aerogear.mobile.security.checks.TestCheck"
	securityMetricName := "TestCheck"
	securityMetricPassed := true

	bigSecurityMetricList := SecurityMetrics{}
	for i := 0; i <= clientIdMaxLength+1; i++ {
		bigSecurityMetricList = append(bigSecurityMetricList, SecurityMetric{Id: &securityMetricId, Passed: &securityMetricPassed})
	}

	validAppMetric := &AppMetric{
		SDKVersion: "1",
		ID:         "some-app-id",
		AppVersion: "1",
	}

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
			ExpectedReason: missingClientIdError,
		},
		{
			Name:           "Metric with no clientId should be invalid",
			Metric:         Metric{Data: &MetricData{App: validAppMetric}},
			Valid:          false,
			ExpectedReason: missingClientIdError,
		},
		{
			Name:           "Metric with long clientId should be invalid",
			Metric:         Metric{ClientId: strings.Join(make([]string, clientIdMaxLength+10), "a"), Data: &MetricData{App: validAppMetric}},
			Valid:          false,
			ExpectedReason: clientIdLengthError,
		},
		{
			Name:           "Metric with no Data should be invalid",
			Metric:         Metric{ClientId: "org.aerogear.metrics.testing"},
			Valid:          false,
			ExpectedReason: missingDataError,
		},
		{
			Name:           "Metric with empty Data should be invalid",
			Metric:         Metric{ClientId: "org.aerogear.metrics.testing", Data: &MetricData{}},
			Valid:          false,
			ExpectedReason: missingDataError,
		},
		{
			Name:           "Metric with ClientId and Some Data should be valid",
			Metric:         Metric{ClientId: "org.aerogear.metrics.testing", Data: &MetricData{App: validAppMetric}},
			Valid:          true,
			ExpectedReason: "",
		},
		{
			Name:           "Metric with bad timestamp should be invalid",
			Metric:         Metric{ClientId: "org.aerogear.metrics.testing", ClientTimestamp: "invalid", Data: &MetricData{App: validAppMetric}},
			Valid:          false,
			ExpectedReason: invalidTimestampError,
		},
		{
			Name:           "Metric with valid timestamp should be valid",
			Metric:         Metric{ClientId: "org.aerogear.metrics.testing", ClientTimestamp: "12345", Data: &MetricData{App: validAppMetric}},
			Valid:          true,
			ExpectedReason: "",
		},
		{
			Name:           "Security Metrics with missing id field should be invalid",
			Metric:         Metric{ClientId: "org.aerogear.metrics.testing", Data: &MetricData{Security: &SecurityMetrics{SecurityMetric{Id: nil, Name: &securityMetricName, Passed: &securityMetricPassed}}}},
			Valid:          false,
			ExpectedReason: fmt.Sprintf(securityMetricMissingIdError, 0),
		},
		{
			Name:           "Security Metrics with missing name field should be invalid",
			Metric:         Metric{ClientId: "org.aerogear.metrics.testing", Data: &MetricData{Security: &SecurityMetrics{SecurityMetric{Id: &securityMetricId, Name: nil, Passed: &securityMetricPassed}}}},
			Valid:          false,
			ExpectedReason: fmt.Sprintf(securityMetricMissingNameError, 0),
		},
		{
			Name:           "Security Metrics with missing passed field should be invalid",
			Metric:         Metric{ClientId: "org.aerogear.metrics.testing", Data: &MetricData{Security: &SecurityMetrics{SecurityMetric{Id: &securityMetricId, Name: &securityMetricName, Passed: nil}}}},
			Valid:          false,
			ExpectedReason: fmt.Sprintf(securityMetricMissingPassedError, 0),
		},
		{
			Name:           "Empty Security Metrics slice should be invalid",
			Metric:         Metric{ClientId: "org.aerogear.metrics.testing", Data: &MetricData{Security: &SecurityMetrics{}}},
			Valid:          false,
			ExpectedReason: securityMetricsEmptyError,
		},
		{
			Name:           "Security Metrics slice with length > max length should be valid",
			Metric:         Metric{ClientId: "org.aerogear.metrics.testing", Data: &MetricData{Security: &bigSecurityMetricList}},
			Valid:          false,
			ExpectedReason: securityMetricsLengthError,
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

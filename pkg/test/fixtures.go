package test

import "github.com/aerogear/aerogear-app-metrics/pkg/mobile"

var securityMetricId = "org.aerogear.mobile.security.checks.TestCheck"
var securityMetricName = "TestCheck"
var securityMetricPassed = true

func GetEmptyMetric() mobile.Metric {
	return mobile.Metric{}
}
func GetNoDataMetric() mobile.Metric {
	return mobile.Metric{
		EventType:       "init",
		ClientTimestamp: "1234",
		ClientId:        "client123",
	}
}

func GetEmptyDataMetric() mobile.Metric {
	m := GetNoDataMetric()
	m.Data = &mobile.MetricData{}
	return m
}

func GetNoClientIdMetric() mobile.Metric {
	m := GetEmptyDataMetric()
	m.ClientId = ""
	return m
}

func GetValidInitMetric() mobile.Metric {
	m := GetEmptyDataMetric()
	m.EventType = "init"
	m.Data = &mobile.MetricData{
		App: &mobile.AppMetric{
			ID:         "deadbeef",
			SDKVersion: "1.2.3",
			AppVersion: "27",
		},
		Device: &mobile.DeviceMetric{
			Platform:        "android",
			PlatformVersion: "19",
		},
	}
	return m
}

func GetNoAppInitMetric() mobile.Metric {
	m := GetValidInitMetric()
	m.Data.App = nil
	return m
}

func GetNoDeviceInitMetric() mobile.Metric {
	m := GetValidInitMetric()
	m.Data.Device = nil
	return m
}

func GetLargeClientIdMetric() mobile.Metric {
	m := GetValidInitMetric()
	m.ClientId = "453de743211112345678900987654312345678908776423567890876543678907654356789076543536789765436789076543256789076543256789654321456789654321456789765432567869765432567896543256789765432145678976543214567897654324567896543214567897654321245678976543256789765432145678965432567896543256786543214567865432"
	return m
}

func GetMetricWithTimestamp() mobile.Metric {
	m := GetValidInitMetric()
	m.ClientId = "withTimestamp"
	m.ClientTimestamp = "123456789"
	return m
}

func GetMetricWithInvalidTimestamp() mobile.Metric {
	m := GetValidInitMetric()
	m.ClientId = "withInvalidTimestamp"
	m.ClientTimestamp = "invalid"
	return m
}

func GetValidSecurityMetric() mobile.Metric {
	m := GetEmptySecurityMetric()
	security := mobile.SecurityMetrics{}
	id := "org.aerogear.mobile.security.checks.DeveloperModeCheck"
	name := "DeveloperModeCheck"
	passed := true
	security = append(security, mobile.SecurityMetric{
		Id:     &id,
		Name:   &name,
		Passed: &passed,
	})
	m.Data.Security = &security
	return m
}

func GetIncompleteSecurityMetric() mobile.Metric {
	m := GetEmptySecurityMetric()
	security := mobile.SecurityMetrics{}
	security = append(security, mobile.SecurityMetric{})
	m.Data.Security = &security
	return m
}

func GetNoIdSecurityMetric() mobile.Metric {
	m := GetEmptySecurityMetric()
	security := mobile.SecurityMetrics{}
	security = append(security, mobile.SecurityMetric{
		Name:   &securityMetricName,
		Passed: &securityMetricPassed,
	})
	m.Data.Security = &security
	return m
}

func GetNoPassedSecurityMetric() mobile.Metric {
	m := GetEmptySecurityMetric()
	security := mobile.SecurityMetrics{}
	security = append(security, mobile.SecurityMetric{
		Id:   &securityMetricId,
		Name: &securityMetricName,
	})
	m.Data.Security = &security
	return m
}
func GetNoNameSecurityMetric() mobile.Metric {
	m := GetEmptySecurityMetric()
	security := mobile.SecurityMetrics{}
	security = append(security, mobile.SecurityMetric{
		Id:     &securityMetricId,
		Passed: &securityMetricPassed,
	})
	m.Data.Security = &security
	return m
}

func GetEmptySecurityMetric() mobile.Metric {
	m := GetValidInitMetric()
	m.EventType = "security"
	m.Data.Security = &mobile.SecurityMetrics{}
	return m
}

func GetOverfilledSecurityMetric() mobile.Metric {
	bigSecurityMetricList := mobile.SecurityMetrics{}
	for i := 0; i <= 129; i++ {
		bigSecurityMetricList = append(bigSecurityMetricList, mobile.SecurityMetric{
			Id:     &securityMetricId,
			Name:   &securityMetricName,
			Passed: &securityMetricPassed,
		})
	}
	m := GetEmptySecurityMetric()
	m.Data.Security = &bigSecurityMetricList
	return m
}

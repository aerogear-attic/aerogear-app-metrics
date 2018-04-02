package test

import "github.com/aerogear/aerogear-app-metrics/pkg/mobile"

func GetEmptyDataMetric() mobile.Metric {
	return mobile.Metric{
		ClientTimestamp: "1234",
		ClientId:        "client123",
		Data:            &mobile.MetricData{},
	}
}

func GetValidInitMetric() mobile.Metric {
	metric := GetEmptyDataMetric()
	metric.Data = &mobile.MetricData{
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
	return metric
}

func GetNoAppInitMetric() mobile.Metric {
	metric := GetValidInitMetric()
	metric.Data.App = nil
	return metric
}

func GetNoDeviceInitMetric() mobile.Metric {
	metric := GetValidInitMetric()
	metric.Data.Device = nil
	return metric
}

func GetLargeClientIdMetric() mobile.Metric {
	metric := GetValidInitMetric()
	metric.ClientId = "453de743211112345678900987654312345678908776423567890876543678907654356789076543536789765436789076543256789076543256789654321456789654321456789765432567869765432567896543256789765432145678976543214567897654324567896543214567897654321245678976543256789765432145678965432567896543256786543214567865432"
	return metric
}

func GetMetricWithTimestamp() mobile.Metric {
	metric := GetValidInitMetric()
	metric.ClientId = "withTimestamp"
	metric.ClientTimestamp = "123456789"
	return metric
}

func GetValidSecurityMetric() mobile.Metric {
	metric := GetValidInitMetric()
	security := mobile.SecurityMetrics{}
	id := "org.aerogear.mobile.security.checks.DeveloperModeCheck"
	name := "DeveloperModeCheck"
	passed := true
	security = append(security, mobile.SecurityMetric{
		Id:     &id,
		Name:   &name,
		Passed: &passed,
	})
	metric.Data.Security = &security
	return metric
}

func GetIncompleteSecurityMetric() mobile.Metric {
	metric := GetValidInitMetric()
	security := mobile.SecurityMetrics{}
	passed := true
	security = append(security, mobile.SecurityMetric{
		Passed: &passed,
	})
	metric.Data.Security = &security
	return metric
}

func GetEmptySecurityMetric() mobile.Metric {
	metric := GetValidInitMetric()
	metric.Data.Security = &mobile.SecurityMetrics{}
	return metric
}

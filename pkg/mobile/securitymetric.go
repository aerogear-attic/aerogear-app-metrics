package mobile

import "fmt"

type SecurityMetric struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Passed bool   `json:"passed"`
}

const securityMetricsMaxLength = 30
const missingSecurityError = "missing data.security in security-type payload"
const securityMetricMissingIdError = "invalid element in data.security at position %v, id must be included"
const securityMetricMissingNameError = "invalid element in data.security at position %v, name must be included"
const securityMetricMissingPassedError = "invalid element in data.security at position %v, passed must be included"

func (sm *SecurityMetric) Validate(i int) (valid bool, reason string) {
	if sm.Id == "" {
		return false, fmt.Sprintf(securityMetricMissingIdError, i)
	}
	if sm.Name == "" {
		return false, fmt.Sprintf(securityMetricMissingNameError, i)
	}
	return true, ""
}

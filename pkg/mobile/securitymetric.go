package mobile

import "fmt"

type SecurityMetric struct {
	Id     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Passed *bool   `json:"passed,omitempty"`
}

const securityMetricsMaxLength = 30
const missingSecurityError = "missing data.security in security-type payload"
const securityMetricMissingIdError = "invalid element in data.security at position %v, id must be included"
const securityMetricMissingNameError = "invalid element in data.security at position %v, name must be included"
const securityMetricMissingPassedError = "invalid element in data.security at position %v, passed must be included"

func (sm *SecurityMetric) Validate(i int) (valid bool, reason string) {
	if sm.Id == nil {
		return false, fmt.Sprintf(securityMetricMissingIdError, i)
	}
	if sm.Name == nil {
		return false, fmt.Sprintf(securityMetricMissingNameError, i)
	}
	if sm.Passed == nil {
		return false, fmt.Sprintf(securityMetricMissingPassedError, i)
	}
	return true, ""
}

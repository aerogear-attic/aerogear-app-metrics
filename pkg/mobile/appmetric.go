package mobile

type AppMetric struct {
	ID         string `json:"appId"`
	SDKVersion string `json:"sdkVersion"`
	AppVersion string `json:"appVersion"`
}

const missingIDError = "missing data.app.appId in payload"
const missingSDKVersionError = "missing data.app.sdkVersion in payload"
const missingAppVersionError = "missing data.app.appVersion in payload"

func (app *AppMetric) Validate() (valid bool, reason string) {
	if app.ID == "" {
		return false, missingIDError
	}
	if app.SDKVersion == "" {
		return false, missingSDKVersionError
	}
	if app.AppVersion == "" {
		return false, missingAppVersionError
	}
	return true, ""
}

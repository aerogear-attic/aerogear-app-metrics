package mobile

type DeviceMetric struct {
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platformVersion"`
}

const missingPlatformError = "missing data.app.appId in payload"
const missingPlatformVersionError = "missing data.app.sdkVersion in payload"

func (dev *DeviceMetric) Validate() (valid bool, reason string) {
	if dev.Platform == "" {
		return false, missingPlatformError
	}
	if dev.PlatformVersion == "" {
		return false, missingPlatformVersionError
	}
	return true, ""
}

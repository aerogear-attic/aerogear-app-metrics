package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aerogear/aerogear-app-metrics/pkg/config"
	"github.com/aerogear/aerogear-app-metrics/pkg/dao"
	"github.com/aerogear/aerogear-app-metrics/pkg/mobile"
)

type SeedOptions struct {
	n                int
	seed             int64
	clients          int
	apps             int
	appVersions      int
	sdkVersions      int
	platformVersions int
	metricsTypes     int
}

var platforms = []string{"android", "ios", "cordova"}

var securityNames = []string{"DeveloperModeCheck", "EmulatorCheck", "DebuggerCheck", "RootedCheck", "ScreenLockCheck"}

var boolGenerator = &genCache{}

const (
	appAndDeviceMetrics = 1 << iota
	securityMetrics     = 1 << iota
)

const clientIdLen = 30
const appIdLen = 20

func main() {
	config := config.GetConfig()

	dbHandler := dao.DatabaseHandler{}

	err := dbHandler.Connect(config.DBConnectionString, config.DBMaxConnections)

	if err != nil {
		panic("failed to connect to sql database : " + err.Error())
	}
	defer dbHandler.DB.Close()

	if err := dbHandler.DoInitialSetup(); err != nil {
		panic("failed to perform database setup : " + err.Error())
	}

	metricsDao := dao.NewMetricsDAO(dbHandler.DB)

	metricsService := mobile.NewMetricsService(metricsDao)

	generateMetrics(metricsService)
}

func generateMetrics(metricsService *mobile.MetricsService, opts *SeedOptions) {
	seedValue := opts.seed
	if seedValue == 0 {
		seedValue = time.Now().UnixNano()
	}
	rand.Seed(seedValue)

	// generate fixtures to be selected from
	clients := makeRandomStrings(opts.clients, clientIdLen)
	appIds := makeRandomStrings(opts.apps, appIdLen)
	appVersions := makeRandomSemvers(opts.appVersions)
	sdkVersions := makeRandomSemvers(opts.sdkVersions)
	platformVersions := makeRandomSemvers(opts.platformVersions)

	securityIds := make([]string, len(securityNames))
	for i := 0; i < len(securityNames); i++ {
		securityIds[i] = "org.aerogear.mobile.security.checks." + securityNames[i]
	}

	metricData := new(mobile.MetricData)
	if (opts.metricsTypes & appAndDeviceMetrics) == appAndDeviceMetrics {
		metricData.App = &mobile.AppMetric{
			ID:         appIds[rand.Intn(opts.apps)],
			AppVersion: appVersions[rand.Intn(opts.appVersions)],
			SDKVersion: sdkVersions[rand.Intn(opts.sdkVersions)],
		}
		metricData.Device = &mobile.DeviceMetric{
			Platform:        platforms[rand.Intn(len(platforms))],
			PlatformVersion: platforms[rand.Intn(len(platforms))],
		}
	}
	if (opts.metricsTypes & securityMetrics) == securityMetrics {
		security := mobile.SecurityMetrics{}
		for i := 0; i < len(securityNames); i++ {
			id := RandStringBytesMaskImpr(10)
			passed := boolGenerator.Bool()
			securityMetric := mobile.SecurityMetric{
				Id:     &id,
				Name:   &securityNames[i],
				Passed: &passed,
			}

			security = append(security, securityMetric)
		}

		metricData.Security = &security
	}
	metric := new(mobile.Metric)
	metric.ClientId = clients[rand.Intn(len(clients))]
	metric.Data = metricData
}

func makeRandomSemvers(n int) []string {
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		arr[i] = fmt.Sprintf("%d.%d.%d", rand.Intn(3), rand.Intn(10), rand.Intn(10))
	}
	return arr
}

func makeRandomStrings(n int, len int) []string {
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		arr[i] = RandStringBytesMaskImpr(len)
	}
	return arr
}

// from https://stackoverflow.com/questions/45030618/generate-a-random-bool-in-go
type genCache struct {
	cache     int64
	remaining int
}

func (b *genCache) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = rand.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

// from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

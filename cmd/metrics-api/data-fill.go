package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/aerogear/aerogear-app-metrics/pkg/config"
	"github.com/aerogear/aerogear-app-metrics/pkg/dao"
	"github.com/aerogear/aerogear-app-metrics/pkg/mobile"
)

type SeedOptions struct {
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
	n := *flag.Int("n", 15000, "Number of records to generate")

	opts := &SeedOptions{}
	flag.IntVar(&opts.apps, "apps", 3, "Number of different apps to generate")
	flag.IntVar(&opts.clients, "clients", 100, "Number of different clients to generate")
	flag.IntVar(&opts.appVersions, "appVersions", 3, "Number of different appVersions to use")
	flag.IntVar(&opts.sdkVersions, "sdkVersions", 3, "Number of different sdkVersions to generate")

	// TODO: make metrics types selectable
	opts.metricsTypes = appAndDeviceMetrics | securityMetrics

	if n == 0 || opts.clients == 0 || opts.apps == 0 || opts.appVersions == 0 || opts.sdkVersions == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	service := initMetricsService()
	for i := 0; i < n; i++ {
		metric := generateMetrics(opts)
		// TODO: add options to send metric via HTTP and print JSON to stdout
		service.Create(*metric)
	}
}

func initMetricsService() *mobile.MetricsService {
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
	return metricsService
}

func generateMetrics(opts *SeedOptions) *mobile.Metric {
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
			PlatformVersion: platformVersions[rand.Intn(len(platformVersions))],
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

	return metric
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

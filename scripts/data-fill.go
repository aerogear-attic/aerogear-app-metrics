// CLI to seed random data to the database used by the metrics service
// use the available flags to determine variance and amount of records to generate
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/aerogear/aerogear-app-metrics/pkg/config"
	"github.com/aerogear/aerogear-app-metrics/pkg/dao"
	"github.com/aerogear/aerogear-app-metrics/pkg/mobile"
	"github.com/aerogear/aerogear-app-metrics/pkg/web"
)

type SeedOptions struct {
	seed               int64
	clients            int
	apps               int
	appVersions        int
	sdkVersions        int
	platformVersions   int
	metricsTypes       int
	securityFailChance float64
	httpTarget         string
}

type SeedData struct {
	clients          []string
	clientsPlatforms map[string]string
	appVersions      []string
	sdkVersions      []string
	platformVersions []string
}

type metricsHTTPService struct {
	hostname string
}

func (h *metricsHTTPService) Create(metric mobile.Metric) (mobile.Metric, error) {
	byteBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(byteBuffer).Encode(metric); err != nil {
		return metric, err
	}
	res, err := http.Post(h.hostname+"/metrics", "application/json", byteBuffer)
	if err != nil {
		return metric, err
	}
	if res.StatusCode != 204 {
		return metric, fmt.Errorf("Non OK status code %v", res.Status)
	}
	return metric, nil
}

var platforms = []string{"android", "ios", "cordova"}

var securityNames = []string{"DeveloperModeCheck", "EmulatorCheck", "DebuggerCheck", "RootedCheck", "ScreenLockCheck"}
var securityIds = []string{
	"org.aerogear.mobile.security.checks.DeveloperModeCheck",
	"org.aerogear.mobile.security.checks.EmulatorCheck",
	"org.aerogear.mobile.security.checks.DebuggerCheck",
	"org.aerogear.mobile.security.checks.RootedCheck",
	"org.aerogear.mobile.security.checks.ScreenLockCheck",
}

const (
	appAndDeviceMetrics = 1 << iota
	securityMetrics     = 1 << iota
)

const clientIdLen = 30
const appIdLen = 20

func main() {
	n := flag.Int("n", 1000, "Number of records to generate")

	opts := &SeedOptions{}
	flag.IntVar(&opts.apps, "apps", 3, "Number of different apps to generate")
	flag.IntVar(&opts.clients, "clients", 100, "Number of different clients to generate")
	flag.IntVar(&opts.appVersions, "appVersions", 3, "Number of different appVersions to use")
	flag.IntVar(&opts.sdkVersions, "sdkVersions", 3, "Number of different sdkVersions to generate")
	flag.IntVar(&opts.platformVersions, "platformVersions", 3, "Number of different platformVersions to generate")
	flag.Float64Var(&opts.securityFailChance, "fail", 0.2, "Float chance of a security check failing randomly, use 0 to always pass")

	flag.StringVar(&opts.httpTarget, "h", "", "address of a running server to target metrics generation via http. i.e. http://localhost:3000")

	flag.Int64Var(&opts.seed, "seed", time.Now().UnixNano(), "Explicit seed value to use for replicable results, defaults to system time")

	// TODO: make metrics types selectable or also random
	opts.metricsTypes = appAndDeviceMetrics | securityMetrics

	flag.Parse()

	if *n == 0 || opts.clients == 0 || opts.apps == 0 || opts.appVersions == 0 || opts.sdkVersions == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	seedValue := opts.seed
	rand.Seed(seedValue)

	// generate fixtures to be selected from
	clients := makeRandomStrings(opts.clients, clientIdLen)
	appVersions := makeRandomSemvers(opts.appVersions)
	sdkVersions := makeRandomSemvers(opts.sdkVersions)
	platformVersions := makeRandomSemvers(opts.platformVersions)
	// associante a client with a specific platform
	clientsPlatforms := make(map[string]string)
	for _, client := range clients {
		clientsPlatforms[client] = platforms[rand.Intn(len(platforms))]
	}

	seedData := &SeedData{
		clients:          clients,
		clientsPlatforms: clientsPlatforms,
		appVersions:      appVersions,
		sdkVersions:      sdkVersions,
		platformVersions: platformVersions,
	}

	service := getMetricsServiceImpl(opts)
	for i := 0; i < *n; i++ {
		metric := generateMetrics(opts, seedData)
		// TODO: add options to send metric via HTTP and print JSON to stdout
		_, err := service.Create(*metric)
		if err != nil {
			log.Printf("Error creating record %d: %v\n", i+1, err)
		} else {
			fmt.Printf("Created record %d\n", i+1)
		}
	}
}

func getMetricsServiceImpl(opts *SeedOptions) web.MetricsServiceInterface {
	if opts.httpTarget != "" {
		return initHttpService(opts.httpTarget)
	}
	return initMetricsService()
}

func initHttpService(host string) web.MetricsServiceInterface {
	fmt.Printf("Utilizing http metrics creation targetting host %v\n", host)
	return &metricsHTTPService{
		hostname: host,
	}
}

func initMetricsService() *mobile.MetricsService {
	fmt.Println("Targetting default postgresql instance")
	config := config.GetConfig()

	dbHandler := dao.DatabaseHandler{}

	err := dbHandler.Connect(config.DBConnectionString, config.DBMaxConnections)

	if err != nil {
		panic("failed to connect to sql database : " + err.Error())
	}

	if err := dbHandler.DoInitialSetup(); err != nil {
		panic("failed to perform database setup : " + err.Error())
	}

	metricsDao := dao.NewMetricsDAO(dbHandler.DB)

	metricsService := mobile.NewMetricsService(metricsDao)
	return metricsService
}

func generateMetrics(opts *SeedOptions, fixtures *SeedData) *mobile.Metric {
	metricData := &mobile.MetricData{}
	client := fixtures.clients[rand.Intn(opts.clients)]

	if (opts.metricsTypes & appAndDeviceMetrics) == appAndDeviceMetrics {
		metricData.App = &mobile.AppMetric{
			ID:         fmt.Sprintf("app%d", rand.Intn(opts.apps)),
			AppVersion: fixtures.appVersions[rand.Intn(opts.appVersions)],
			SDKVersion: fixtures.sdkVersions[rand.Intn(opts.sdkVersions)],
		}
		metricData.Device = &mobile.DeviceMetric{
			Platform:        fixtures.clientsPlatforms[client],
			PlatformVersion: fixtures.platformVersions[rand.Intn(opts.platformVersions)],
		}
	}
	if (opts.metricsTypes & securityMetrics) == securityMetrics {
		security := mobile.SecurityMetrics{}
		for i := 0; i < len(securityNames); i++ {
			passed := true
			if opts.securityFailChance > 0.0 {
				passed = rand.Float64() > opts.securityFailChance
			}
			securityMetric := mobile.SecurityMetric{
				Id:     &securityIds[i],
				Name:   &securityNames[i],
				Passed: &passed,
			}

			security = append(security, securityMetric)
		}

		metricData.Security = &security
	}

	return &mobile.Metric{
		ClientId: client,
		Data:     metricData,
	}
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

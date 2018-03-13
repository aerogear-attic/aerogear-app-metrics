package dao

import (
	"log"
	"time"

	"github.com/aerogear/aerogear-app-metrics/pkg/mobile"
	client "github.com/influxdata/influxdb/client/v2"
)

type MetricsDAO struct {
}

// Create a metrics record
func (m *MetricsDAO) Create(clientId string, metric mobile.Metric, clientTime *time.Time) error {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "mydb",
		Precision: "s",
	})
	if err != nil {
		return err
	}

	if metric.Data.App != nil {
		appMetric := metric.Data.App
		tags := map[string]string{
			"clientId": clientId,
			"appId":    appMetric.ID,
		}
		fields := map[string]interface{}{
			"sdk_version": appMetric.SDKVersion,
			"app_version": appMetric.AppVersion,
		}

		pt, err := client.NewPoint("app", tags, fields, time.Now())
		if err != nil {
			return err
		}
		bp.AddPoint(pt)
	}
	if metric.Data.Device != nil {
		deviceMetric := metric.Data.Device
		tags := map[string]string{
			"clientId": clientId,
			"platform": deviceMetric.Platform,
		}
		fields := map[string]interface{}{
			"platform_version": deviceMetric.PlatformVersion,
		}

		pt, err := client.NewPoint("platform", tags, fields, time.Now())
		if err != nil {
			return err
		}
		bp.AddPoint(pt)
	}
	if metric.Data.Security != nil {
		for _, securityMetric := range *metric.Data.Security {
			tags := map[string]string{
				"clientId": clientId,
				"id":       securityMetric.Id,
				"name":     securityMetric.Name,
			}
			fields := map[string]interface{}{
				"passed": securityMetric.Passed,
			}

			pt, err := client.NewPoint("security", tags, fields, time.Now())
			if err != nil {
				return err
			}
			bp.AddPoint(pt)
		}
	}
	// Write the batch
	err = c.Write(bp)
	return err
}

// IsHealthy checks that we are connected to the database
// This will be used by the healthcheck
func (m *MetricsDAO) IsHealthy() error {
	return nil
}

func NewMetricsDAO() *MetricsDAO {
	return &MetricsDAO{}
}

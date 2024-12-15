package service

import (
	"context"
	"sensor_monitoring_be/models"
)

type ReportService interface {
	CreateReport(ctx context.Context, data models.GeneratedReport) (models.GeneratedReport, error)
	ReportList(ctx context.Context) ([]map[string]interface{}, error)
	GetSensor(ctx context.Context, deviceId int64) ([]map[string]interface{}, error)
	GetSensorData(ctx context.Context, sensorId int64, start_date string, end_date string) ([]map[string]interface{}, error)
	UpdateReport(ctx context.Context, deviceId int64, startDate string, endDate string, status string, file string) error
	DownloadReportFile(ctx context.Context, id int64) (string, error)
}

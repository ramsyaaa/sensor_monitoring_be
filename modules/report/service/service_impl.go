package service

import (
	"context"

	"sensor_monitoring_be/models"
	"sensor_monitoring_be/modules/report/repository"
)

type service struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) ReportService {
	return &service{repo: repo}
}

func (s *service) CreateReport(ctx context.Context, data models.GeneratedReport) (models.GeneratedReport, error) {
	return s.repo.CreateReport(ctx, data)
}

func (s *service) ReportList(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.ReportList(ctx)
}

func (s *service) GetSensor(ctx context.Context, deviceId int64) ([]map[string]interface{}, error) {
	return s.repo.GetSensor(ctx, deviceId)
}

func (s *service) GetSensorData(ctx context.Context, sensorId int64, start_date string, end_date string) ([]map[string]interface{}, error) {
	return s.repo.GetSensorData(ctx, sensorId, start_date, end_date)
}

func (s *service) UpdateReport(ctx context.Context, deviceId int64, startDate string, endDate string, status string, file string) error {
	return s.repo.UpdateReport(ctx, deviceId, startDate, endDate, status, file)
}
func (s *service) DownloadReportFile(ctx context.Context, id int64) (string, error) {
	return s.repo.DownloadReportFile(ctx, id)
}

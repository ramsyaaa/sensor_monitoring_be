package service

import (
	"context"

	"sensor_monitoring_be/modules/geomapping/repository"
)

type service struct {
	repo repository.GeoMappingRepository
}

func NewGeoMappingService(repo repository.GeoMappingRepository) GeoMappingService {
	return &service{repo: repo}
}

func (s *service) GetDevice(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetDevice(ctx)
}

func (s *service) GetSensor(ctx context.Context, deviceId int) ([]map[string]interface{}, error) {
	return s.repo.GetSensor(ctx, deviceId)
}
func (s *service) UpdateSensorData(ctx context.Context, sensorId int, data map[string]interface{}) error {
	return s.repo.UpdateSensorData(ctx, sensorId, data)
}

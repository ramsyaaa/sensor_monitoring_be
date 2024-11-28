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

func (s *service) GetDevice(ctx context.Context, group_id, city_id, district_id, subdistrict_id int, keyword string) ([]map[string]interface{}, error) {
	return s.repo.GetDevice(ctx, group_id, city_id, district_id, subdistrict_id, keyword)
}

func (s *service) GetDeviceDetail(ctx context.Context, deviceId int) ([]map[string]interface{}, error) {
	return s.repo.GetDeviceDetail(ctx, deviceId)
}

func (s *service) GetSensor(ctx context.Context, deviceId int) ([]map[string]interface{}, error) {
	return s.repo.GetSensor(ctx, deviceId)
}

func (s *service) UpdateSensorData(ctx context.Context, sensorId int, data map[string]interface{}) error {
	return s.repo.UpdateSensorData(ctx, sensorId, data)
}

func (s *service) UpdateDeviceData(ctx context.Context, deviceId int, data map[string]interface{}) error {
	return s.repo.UpdateDeviceData(ctx, deviceId, data)
}

func (s *service) GetGroup(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetGroup(ctx)
}

func (s *service) GetCity(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetCity(ctx)
}

func (s *service) GetDistrict(ctx context.Context, cityId int) ([]map[string]interface{}, error) {
	return s.repo.GetDistrict(ctx, cityId)
}

func (s *service) GetSubDistrict(ctx context.Context, districtId int) ([]map[string]interface{}, error) {
	return s.repo.GetSubDistrict(ctx, districtId)
}

func (s *service) Dashboard(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.Dashboard(ctx)
}

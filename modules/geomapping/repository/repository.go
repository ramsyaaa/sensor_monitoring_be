package repository

import (
	"context"
)

type GeoMappingRepository interface {
	GetDevice(ctx context.Context) ([]map[string]interface{}, error)
	GetSensor(ctx context.Context, deviceID int) ([]map[string]interface{}, error)
	UpdateSensorData(ctx context.Context, sensorId int, data map[string]interface{}) error
	UpdateDeviceData(ctx context.Context, deviceId int, data map[string]interface{}) error
	GetGroup(ctx context.Context) ([]map[string]interface{}, error)
	GetCity(ctx context.Context) ([]map[string]interface{}, error)
	GetDistrict(ctx context.Context, cityID int) ([]map[string]interface{}, error)
	GetSubDistrict(ctx context.Context, districtID int) ([]map[string]interface{}, error)
}

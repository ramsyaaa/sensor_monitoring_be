package service

import (
	"context"
)

type GeoMappingService interface {
	GetDevice(ctx context.Context) ([]map[string]interface{}, error)
	GetSensor(ctx context.Context, deviceId int) ([]map[string]interface{}, error)
	UpdateSensorData(ctx context.Context, sensorId int, data map[string]interface{}) error
}

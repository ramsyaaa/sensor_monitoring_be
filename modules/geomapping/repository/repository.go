package repository

import (
	"context"
)

type GeoMappingRepository interface {
	GetDevice(ctx context.Context) ([]map[string]interface{}, error)
	GetSensor(ctx context.Context, deviceID int) ([]map[string]interface{}, error)
	UpdateSensorData(ctx context.Context, sensorId int, data map[string]interface{}) error
}

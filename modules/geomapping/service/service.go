package service

import (
	"context"
)

type GeoMappingService interface {
	GetDevice(ctx context.Context, groupID, cityID, districtID, subdistrictID int) ([]map[string]interface{}, error)
	GetDeviceDetail(ctx context.Context, deviceId int) ([]map[string]interface{}, error)
	GetSensor(ctx context.Context, deviceId int) ([]map[string]interface{}, error)
	UpdateSensorData(ctx context.Context, sensorId int, data map[string]interface{}) error
	UpdateDeviceData(ctx context.Context, deviceId int, data map[string]interface{}) error
	GetGroup(ctx context.Context) ([]map[string]interface{}, error)
	GetCity(ctx context.Context) ([]map[string]interface{}, error)
	GetDistrict(ctx context.Context, cityId int) ([]map[string]interface{}, error)
	GetSubDistrict(ctx context.Context, districtId int) ([]map[string]interface{}, error)
}

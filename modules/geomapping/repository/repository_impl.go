package repository

import (
	"context"
	"errors"

	"sensor_monitoring_be/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewGeoMappingRepository(db *gorm.DB) GeoMappingRepository {
	return &repository{db: db}
}

func (r *repository) GetDevice(ctx context.Context) ([]map[string]interface{}, error) {
	var devices []map[string]interface{}
	err := r.db.WithContext(ctx).Model(&models.Device{}).Select("id, device_name, device_no, lat, lng").Find(&devices).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return devices, err
}

func (r *repository) GetSensor(ctx context.Context, deviceId int) ([]map[string]interface{}, error) {
	var sensors []map[string]interface{}
	err := r.db.WithContext(ctx).Model(&models.Sensor{}).Where("device_id = ?", deviceId).Select("id, device_id, sensor_name, lat, lng").Find(&sensors).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return sensors, err
}

func (r *repository) UpdateSensorData(ctx context.Context, sensorId int, data map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(&models.Sensor{}).Where("id = ?", sensorId).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

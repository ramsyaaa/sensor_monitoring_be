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

func (r *repository) UpdateDeviceData(ctx context.Context, deviceId int, data map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(&models.Device{}).Where("id = ?", deviceId).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetGroup(ctx context.Context) ([]map[string]interface{}, error) {
	var groups []map[string]interface{}
	err := r.db.WithContext(ctx).Model(&models.Group{}).Select("group_id, group_name").Find(&groups).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return groups, err
}

func (r *repository) GetCity(ctx context.Context) ([]map[string]interface{}, error) {
	var cities []map[string]interface{}
	err := r.db.WithContext(ctx).Model(&models.City{}).Select("city_id, city_name").Find(&cities).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return cities, err
}

func (r *repository) GetDistrict(ctx context.Context, cityID int) ([]map[string]interface{}, error) {
	var districts []map[string]interface{}
	err := r.db.WithContext(ctx).Model(&models.District{}).Where("city_id = ?", cityID).Select("district_id, district_name").Find(&districts).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return districts, err
}

func (r *repository) GetSubDistrict(ctx context.Context, districtID int) ([]map[string]interface{}, error) {
	var subDistricts []map[string]interface{}
	err := r.db.WithContext(ctx).Model(&models.Subdistrict{}).Where("district_id = ?", districtID).Select("subdistrict_id, subdistrict_name").Find(&subDistricts).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return subDistricts, err
}

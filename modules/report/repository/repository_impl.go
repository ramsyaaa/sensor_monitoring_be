package repository

import (
	"context"
	"errors"
	"time"

	"sensor_monitoring_be/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &repository{db: db}
}

func (r *repository) CreateReport(ctx context.Context, data models.GeneratedReport) (models.GeneratedReport, error) {
	err := r.db.WithContext(ctx).Model(&models.GeneratedReport{}).Create(&data).Error
	return data, err
}

func (r *repository) ReportList(ctx context.Context) ([]map[string]interface{}, error) {
	var reports []map[string]interface{}
	err := r.db.WithContext(ctx).Model(&models.GeneratedReport{}).Find(&reports).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return reports, err
}

func (r *repository) GetSensor(ctx context.Context, deviceId int64) ([]map[string]interface{}, error) {
	var sensors []map[string]interface{}
	err := r.db.WithContext(ctx).Model(&models.Sensor{}).Where("device_id = ?", deviceId).Select("id, device_id, sensor_name, is_line, lat, lng").Find(&sensors).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return sensors, err
}

func (r *repository) GetSensorData(ctx context.Context, sensorId int64, start_date string, end_date string) ([]map[string]interface{}, error) {
	var sensorData []map[string]interface{}
	err := r.db.WithContext(ctx).Raw("SELECT sensor_data.*, sensors.sensor_name FROM sensor_data JOIN sensors ON sensor_data.sensor_id = sensors.id WHERE sensor_data.sensor_id = ? AND sensor_data.date >= ? AND sensor_data.date <= ?", sensorId, start_date, end_date).Scan(&sensorData).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return sensorData, err
}

func (r *repository) UpdateReport(ctx context.Context, deviceId int64, startDate string, endDate string, status string, file string) error {
	err := r.db.WithContext(ctx).Model(&models.GeneratedReport{}).Where("device_id = ?", deviceId).Where("start_date = ?", startDate).Where("end_date = ?", endDate).Updates(map[string]interface{}{
		"status":       status,
		"file":         file,
		"generated_at": time.Now(),
	}).Error
	return err
}

func (r *repository) DownloadReportFile(ctx context.Context, id int64) (string, error) {
	var file string
	err := r.db.WithContext(ctx).Model(&models.GeneratedReport{}).Where("id = ?", id).Pluck("file", &file).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}
		return "", err
	}
	return file, nil
}

package models

import "time"

func (SensorData) TableName() string {
	return "sensor_data"
}

type SensorData struct {
	ID       int64 `gorm:"primaryKey;autoIncrement"`
	DeviceID int64
	SensorID int64
	Date     time.Time `gorm:"autoCreateTime"`
	Value    float64
}

package models

import "time"

func (GeneratedReport) TableName() string {
	return "generated_reports"
}

type GeneratedReport struct {
	ID          int       `json:"id"`
	DeviceID    int       `json:"device_id"`
	Status      string    `json:"status" gorm:"default:'pending'"`
	File        string    `json:"file"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	GeneratedAt time.Time `json:"generated_at" gorm:"null"`
}

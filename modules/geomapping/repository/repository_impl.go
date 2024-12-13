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

func (r *repository) GetDevice(ctx context.Context, groupId, cityId, districtId, subdistrictId int, keyword string) ([]map[string]interface{}, error) {
	var devices []map[string]interface{}

	query := `SELECT d.id, d.device_name, d.is_line, d.device_no, d.lat, d.lng, 
		d.city_id, d.group_id, d.district_id, d.subdistrict_id, 
		d.point_code, d.address, d.electrical_panel, d.surrounding_waters, 
		d.location_information, d.note,
		g.group_name, c.city_name, dt.district_name, sd.subdistrict_name,
		-- Separate columns for the three sensors
		MAX(CASE WHEN s.sensor_name = 'Water level' THEN s.id END) AS water_level_sensor_id,
		MAX(CASE WHEN s.sensor_name = 'Water level' THEN s.sensor_name END) AS water_level_sensor_name,
		MAX(CASE WHEN s.sensor_name = 'Water level' THEN s.value END) AS water_level_value,
		MAX(CASE WHEN s.sensor_name = 'Water level' THEN s.unit END) AS water_level_unit,
		MAX(CASE WHEN s.sensor_name = 'Flow velocity' THEN s.id END) AS flow_velocity_sensor_id,
		MAX(CASE WHEN s.sensor_name = 'Flow velocity' THEN s.sensor_name END) AS flow_velocity_sensor_name,
		MAX(CASE WHEN s.sensor_name = 'Flow velocity' THEN s.value END) AS flow_velocity_value,
		MAX(CASE WHEN s.sensor_name = 'Flow velocity' THEN s.unit END) AS flow_velocity_unit,
		MAX(CASE WHEN s.sensor_name = 'Instantaneous flow' THEN s.id END) AS instantaneous_flow_sensor_id,
		MAX(CASE WHEN s.sensor_name = 'Instantaneous flow' THEN s.sensor_name END) AS instantaneous_flow_sensor_name,
		MAX(CASE WHEN s.sensor_name = 'Instantaneous flow' THEN s.value END) AS instantaneous_flow_value,
		MAX(CASE WHEN s.sensor_name = 'Instantaneous flow' THEN s.unit END) AS instantaneous_flow_unit
	FROM devices d
	LEFT JOIN groups g ON d.group_id = g.group_id
	LEFT JOIN cities c ON d.city_id = c.city_id 
	LEFT JOIN districts dt ON d.district_id = dt.district_id
	LEFT JOIN subdistricts sd ON d.subdistrict_id = sd.subdistrict_id
	LEFT JOIN sensors s ON d.id = s.device_id AND s.sensor_name IN ('Water level', 'Flow velocity', 'Instantaneous flow')
	WHERE true`

	var params []interface{}

	if groupId != 0 {
		query += ` AND d.group_id = ?`
		params = append(params, groupId)
	}
	if cityId != 0 {
		query += ` AND d.city_id = ?`
		params = append(params, cityId)
	}
	if districtId != 0 {
		query += ` AND d.district_id = ?`
		params = append(params, districtId)
	}
	if subdistrictId != 0 {
		query += ` AND d.subdistrict_id = ?`
		params = append(params, subdistrictId)
	}
	if keyword != "" {
		query += ` AND (LOWER(d.address) LIKE LOWER(?) OR LOWER(d.point_code) LIKE LOWER(?) OR LOWER(g.group_name) LIKE LOWER(?) OR LOWER(c.city_name) LIKE LOWER(?) OR LOWER(dt.district_name) LIKE LOWER(?) OR LOWER(sd.subdistrict_name) LIKE LOWER(?))`
		keyword = "%" + keyword + "%"
		params = append(params, keyword, keyword, keyword, keyword, keyword, keyword)
	}

	query += ` GROUP BY d.id, d.device_name, d.is_line, d.device_no, d.lat, d.lng, 
		d.city_id, d.group_id, d.district_id, d.subdistrict_id,
		d.point_code, d.address, d.electrical_panel, d.surrounding_waters,
		d.location_information, d.note,
		g.group_name, c.city_name, dt.district_name, sd.subdistrict_name
		ORDER BY d.is_line DESC`

	err := r.db.WithContext(ctx).Raw(query, params...).Scan(&devices).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return devices, err
}

func (r *repository) GetDeviceDetail(ctx context.Context, deviceId int) ([]map[string]interface{}, error) {
	var device []map[string]interface{}
	err := r.db.WithContext(ctx).Model(&models.Device{}).Where("id = ?", deviceId).Select("id, device_name, device_no, lat, lng, city_id, group_id, district_id, subdistrict_id, point_code, address, electrical_panel, surrounding_waters, location_information, note").Find(&device).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return device, err
}

func (r *repository) GetSensor(ctx context.Context, deviceId int) ([]map[string]interface{}, error) {
	var sensors []map[string]interface{}
	err := r.db.WithContext(ctx).Model(&models.Sensor{}).Where("device_id = ?", deviceId).Select("id, device_id, sensor_name, is_line, lat, lng").Find(&sensors).Error
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

func (r *repository) Dashboard(ctx context.Context) ([]map[string]interface{}, error) {
	var devices []map[string]interface{}

	query := `SELECT 
		g.group_name,
		d.id,
		d.device_name, 
		d.is_line,
		d.device_no,
		d.lat,
		d.lng,
		d.city_id,
		d.group_id,
		d.district_id,
		d.subdistrict_id,
		d.point_code,
		d.address,
		d.electrical_panel,
		d.surrounding_waters,
		d.location_information,
		d.note,
		c.city_name,
		dt.district_name,
		sd.subdistrict_name,
		-- Separate columns for the three sensors
		MAX(CASE WHEN s.sensor_name = 'Water level' THEN s.id END) AS water_level_sensor_id,
		MAX(CASE WHEN s.sensor_name = 'Water level' THEN s.sensor_name END) AS water_level_sensor_name,
		MAX(CASE WHEN s.sensor_name = 'Water level' THEN s.value END) AS water_level_value,
		MAX(CASE WHEN s.sensor_name = 'Water level' THEN s.unit END) AS water_level_unit,
		MAX(CASE WHEN s.sensor_name = 'Flow velocity' THEN s.id END) AS flow_velocity_sensor_id,
		MAX(CASE WHEN s.sensor_name = 'Flow velocity' THEN s.sensor_name END) AS flow_velocity_sensor_name,
		MAX(CASE WHEN s.sensor_name = 'Flow velocity' THEN s.value END) AS flow_velocity_value,
		MAX(CASE WHEN s.sensor_name = 'Flow velocity' THEN s.unit END) AS flow_velocity_unit,
		MAX(CASE WHEN s.sensor_name = 'Instantaneous flow' THEN s.id END) AS instantaneous_flow_sensor_id,
		MAX(CASE WHEN s.sensor_name = 'Instantaneous flow' THEN s.sensor_name END) AS instantaneous_flow_sensor_name,
		MAX(CASE WHEN s.sensor_name = 'Instantaneous flow' THEN s.value END) AS instantaneous_flow_value,
		MAX(CASE WHEN s.sensor_name = 'Instantaneous flow' THEN s.unit END) AS instantaneous_flow_unit
	FROM devices d
	LEFT JOIN groups g ON d.group_id = g.group_id
	LEFT JOIN cities c ON d.city_id = c.city_id 
	LEFT JOIN districts dt ON d.district_id = dt.district_id
	LEFT JOIN subdistricts sd ON d.subdistrict_id = sd.subdistrict_id
	LEFT JOIN sensors s ON d.id = s.device_id AND s.sensor_name IN ('Water level', 'Flow velocity', 'Instantaneous flow')
	GROUP BY g.group_name, d.id, d.device_name, d.is_line, d.device_no, d.lat, d.lng, d.city_id, d.group_id, d.district_id, d.subdistrict_id, d.point_code, d.address, d.electrical_panel, d.surrounding_waters, d.location_information, d.note, c.city_name, dt.district_name, sd.subdistrict_name
	ORDER BY g.group_name ASC`

	err := r.db.WithContext(ctx).Raw(query).Scan(&devices).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Track order of group names
	var groupOrder []string
	seenGroups := make(map[string]bool)

	// Group devices by group_name while preserving order
	groupedDevices := make(map[string][]map[string]interface{})
	for _, device := range devices {
		groupName, ok := device["group_name"].(string)
		if !ok || groupName == "" {
			continue
		}
		if !seenGroups[groupName] {
			groupOrder = append(groupOrder, groupName)
			seenGroups[groupName] = true
		}
		groupedDevices[groupName] = append(groupedDevices[groupName], device)
	}

	// Convert to final format preserving order
	result := make([]map[string]interface{}, 0)
	for _, groupName := range groupOrder {
		groupMap := map[string]interface{}{
			groupName: groupedDevices[groupName],
		}
		result = append(result, groupMap)
	}

	return result, err
}

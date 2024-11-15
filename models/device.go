package models

import "time"

func (Device) TableName() string {
	return "devices"
}

type Device struct {
	ID                  int64     `json:"id"`
	DeviceName          string    `json:"deviceName"`
	DeviceNo            string    `json:"deviceNo"`
	GroupId             int       `json:"group_id"`
	CityId              int       `json:"city_id"`
	DistrictId          int       `json:"district_id"`
	SubdistrictId       int       `json:"subdistrict_id"`
	IsDelete            int       `json:"isDelete"`
	IsLine              int       `json:"isLine"`
	Lat                 string    `json:"lat"`
	Linktype            string    `json:"linktype"`
	Lng                 string    `json:"lng"`
	ParentUser          int       `json:"parentUser"`
	UserId              int       `json:"userId"`
	UserName            string    `json:"userName"`
	CreateDate          time.Time `json:"createDate"`
	DefaultTimescale    string    `json:"defaultTimescale"`
	IocUrl              string    `json:"iocUrl"`
	IotDeviceId         string    `json:"iotDeviceId"`
	IsAlarms            string    `json:"isAlarms"`
	ProductId           string    `json:"productId"`
	ProductType         string    `json:"productType"`
	ProtocolLabel       string    `json:"protocolLabel"`
	Remark              string    `json:"remark"`
	TimeZone            string    `json:"time_zone"`
	PointCode           string    `json:"point_code"`
	Address             string    `json:"address"`
	ElectricalPanel     string    `json:"electrical_panel"`
	SurroundingWaters   string    `json:"surrounding_waters"`
	LocationInformation string    `json:"location_information"`
	Note                string    `json:"note"`
}

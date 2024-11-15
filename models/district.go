package models

func (District) TableName() string {
	return "districts"
}

type District struct {
	DistrictId   int    `json:"district_id"`
	CityId       int    `json:"city_id"`
	DistrictName string `json:"district_name"`
}

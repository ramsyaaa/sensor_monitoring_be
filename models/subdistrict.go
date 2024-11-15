package models

func (Subdistrict) TableName() string {
	return "subdistricts"
}

type Subdistrict struct {
	SubdistrictId   int    `json:"subdistrict_id"`
	DistrictId      int    `json:"district_id"`
	SubdistrictName string `json:"subdistrict_name"`
}

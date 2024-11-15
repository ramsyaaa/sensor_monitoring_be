package models

func (City) TableName() string {
	return "cities"
}

type City struct {
	CityId   int    `json:"city_id"`
	CityName string `json:"city_name"`
}

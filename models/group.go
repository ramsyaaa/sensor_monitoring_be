package models

func (Group) TableName() string {
	return "groups"
}

type Group struct {
	GroupId   int    `json:"group_id"`
	GroupName string `json:"group_name"`
}

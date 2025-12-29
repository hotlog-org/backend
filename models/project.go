package models

type Project struct {
	ID     string  `json:"id" gorm:"column:id;type:text;primaryKey"`
	Name   string  `json:"name" gorm:"column:name;type:text"`
	Icon   *string `json:"icon,omitempty" gorm:"column:icon;type:text"`
	Colour *string `json:"colour,omitempty" gorm:"column:colour;type:text"`
}

func (Project) TableName() string {
	return "project"
}

package models

import "gorm.io/datatypes"

type Component struct {
	ID          string         `json:"id" gorm:"column:id;type:text;primaryKey"`
	LayoutID    string         `json:"layoutId" gorm:"column:layout_id;type:text"`
	SchemaID    string         `json:"schemaId" gorm:"column:schema_id;type:text"`
	Type        ComponentType  `json:"type" gorm:"column:type;type:text"`
	Title       string         `json:"title" gorm:"column:title;type:text"`
	Description *string        `json:"description,omitempty" gorm:"column:description;type:text"`
	Inputs      datatypes.JSON `json:"inputs" gorm:"column:inputs;type:jsonb"`
	Index       int            `json:"index" gorm:"column:index"`
}

func (Component) TableName() string {
	return "component"
}

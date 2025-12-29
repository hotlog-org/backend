package models

import "gorm.io/datatypes"

type Event struct {
	ID        string         `json:"id" gorm:"column:id;type:text;primaryKey"`
	ProjectID string         `json:"projectId" gorm:"column:project_id;type:text"`
	SchemaID  string         `json:"schemaId" gorm:"column:schema_id;type:text"`
	Value     datatypes.JSON `json:"value" gorm:"column:value;type:jsonb"`
}

func (Event) TableName() string {
	return "event"
}

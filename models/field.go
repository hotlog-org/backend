package models

import "gorm.io/datatypes"

type Field struct {
	ID       string         `json:"id" gorm:"column:id;type:text;primaryKey"`
	SchemaID string         `json:"schemaId" gorm:"column:schema_id;type:text"`
	ObjectID *string        `json:"objectId,omitempty" gorm:"column:object_id;type:text"`
	Label    string         `json:"label" gorm:"column:label;type:text"`
	Type     FieldType      `json:"type" gorm:"column:type;type:text"`
	Status   string         `json:"status" gorm:"column:status;type:text"`
	Options  datatypes.JSON `json:"options" gorm:"column:options;type:jsonb"`
}

func (Field) TableName() string {
	return "field"
}

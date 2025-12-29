package models

type SchemaObject struct {
	ID       string `json:"id" gorm:"column:id;type:text;primaryKey"`
	SchemaID string `json:"schemaId" gorm:"column:schema_id;type:text"`
}

func (SchemaObject) TableName() string {
	return "object"
}

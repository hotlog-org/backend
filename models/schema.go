package models

type Schema struct {
	ID        string `json:"id" gorm:"column:id;type:text;primaryKey"`
	ProjectID string `json:"projectId" gorm:"column:project_id;type:text"`
}

func (Schema) TableName() string {
	return "schema"
}

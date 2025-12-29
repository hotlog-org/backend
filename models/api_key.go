package models

type APIKey struct {
	ID        string `json:"id" gorm:"column:id;type:text;primaryKey"`
	ProjectID string `json:"projectId" gorm:"column:project_id;type:text"`
	Key       string `json:"key" gorm:"column:key;type:text"`
}

func (APIKey) TableName() string {
	return "api_key"
}

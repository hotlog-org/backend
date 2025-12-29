package models

type Role struct {
	ID        string `json:"id" gorm:"column:id;type:text;primaryKey"`
	ProjectID string `json:"projectId" gorm:"column:project_id;type:text"`
	Name      string `json:"name" gorm:"column:name;type:text"`
}

func (Role) TableName() string {
	return "role"
}

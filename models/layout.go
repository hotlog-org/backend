package models

type Layout struct {
	ID          string  `json:"id" gorm:"column:id;type:text;primaryKey"`
	ProjectID   string  `json:"projectId" gorm:"column:project_id;type:text"`
	Label       string  `json:"label" gorm:"column:label;type:text"`
	Description *string `json:"description,omitempty" gorm:"column:description;type:text"`
}

func (Layout) TableName() string {
	return "layout"
}

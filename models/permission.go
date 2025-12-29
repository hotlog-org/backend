package models

type Permission struct {
	ID         string `json:"id" gorm:"column:id;type:text;primaryKey"`
	RoleID     string `json:"roleId" gorm:"column:role_id;type:text"`
	Permission string `json:"permission" gorm:"column:permission;type:text"`
}

func (Permission) TableName() string {
	return "permission"
}

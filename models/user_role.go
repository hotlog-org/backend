package models

type UserRole struct {
	ID     string `json:"id" gorm:"column:id;type:text;primaryKey"`
	UserID string `json:"userId" gorm:"column:user_id;type:text"`
	RoleID string `json:"roleId" gorm:"column:role_id;type:text"`
}

func (UserRole) TableName() string {
	return "user_role"
}

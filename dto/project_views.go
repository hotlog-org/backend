package dto

type ProjectUserRow struct {
	Name          string `gorm:"column:name"`
	Email         string `gorm:"column:email"`
	EmailVerified bool   `gorm:"column:email_verified"`
	RoleName      string `gorm:"column:role_name"`
}

type ProjectRoleRow struct {
	ID   string `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

type RolePermissionRow struct {
	RoleID     string `gorm:"column:role_id"`
	Permission string `gorm:"column:permission"`
}

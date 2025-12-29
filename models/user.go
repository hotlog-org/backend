package models

import "time"

type User struct {
	ID            string    `json:"id" gorm:"column:id;type:text;primaryKey"`
	Name          string    `json:"name" gorm:"column:name;type:text;not null"`
	Email         string    `json:"email" gorm:"column:email;type:text;not null"`
	EmailVerified bool      `json:"emailVerified" gorm:"column:emailVerified;type:boolean;not null"`
	Image         *string   `json:"image,omitempty" gorm:"column:image;type:text"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

func (User) TableName() string {
	return "user"
}

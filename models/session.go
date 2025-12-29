package models

import "time"

type Session struct {
	ID        string    `json:"id" gorm:"column:id;type:text;primaryKey"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"column:expiresAt;not null"`
	Token     string    `json:"token" gorm:"column:token;type:text;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
	IPAddress *string   `json:"ipAddress,omitempty" gorm:"column:ipAddress;type:text"`
	UserAgent *string   `json:"userAgent,omitempty" gorm:"column:userAgent;type:text"`
	UserID    string    `json:"userId" gorm:"column:userId;type:text;not null"`
}

func (Session) TableName() string {
	return "session"
}

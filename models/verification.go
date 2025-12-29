package models

import "time"

type Verification struct {
	ID         string    `json:"id" gorm:"column:id;type:text;primaryKey"`
	Identifier string    `json:"identifier" gorm:"column:identifier;type:text;not null"`
	Value      string    `json:"value" gorm:"column:value;type:text;not null"`
	ExpiresAt  time.Time `json:"expiresAt" gorm:"column:expiresAt;not null"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

func (Verification) TableName() string {
	return "verification"
}

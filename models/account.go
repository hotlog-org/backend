package models

import "time"

type Account struct {
	ID                    string     `json:"id" gorm:"column:id;type:text;primaryKey"`
	AccountID             string     `json:"accountId" gorm:"column:accountId;type:text;not null"`
	ProviderID            string     `json:"providerId" gorm:"column:providerId;type:text;not null"`
	UserID                string     `json:"userId" gorm:"column:userId;type:text;not null"`
	AccessToken           *string    `json:"accessToken,omitempty" gorm:"column:accessToken;type:text"`
	RefreshToken          *string    `json:"refreshToken,omitempty" gorm:"column:refreshToken;type:text"`
	IDToken               *string    `json:"idToken,omitempty" gorm:"column:idToken;type:text"`
	AccessTokenExpiresAt  *time.Time `json:"accessTokenExpiresAt,omitempty" gorm:"column:accessTokenExpiresAt"`
	RefreshTokenExpiresAt *time.Time `json:"refreshTokenExpiresAt,omitempty" gorm:"column:refreshTokenExpiresAt"`
	Scope                 *string    `json:"scope,omitempty" gorm:"column:scope;type:text"`
	Password              *string    `json:"password,omitempty" gorm:"column:password;type:text"`
	CreatedAt             time.Time  `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt             time.Time  `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

func (Account) TableName() string {
	return "account"
}

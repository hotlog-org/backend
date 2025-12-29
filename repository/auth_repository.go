package repository

import (
	"hotlog.org/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) FindUserByID(id string) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *AuthRepository) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *AuthRepository) FindSessionByToken(token string) (models.Session, error) {
	var session models.Session
	err := r.db.Where("token = ?", token).First(&session).Error
	return session, err
}

func (r *AuthRepository) FindAccount(providerID, accountID string) (models.Account, error) {
	var account models.Account
	err := r.db.Where("providerId = ? AND accountId = ?", providerID, accountID).First(&account).Error
	return account, err
}

func (r *AuthRepository) FindVerification(identifier, value string) (models.Verification, error) {
	var verification models.Verification
	err := r.db.Where("identifier = ? AND value = ?", identifier, value).First(&verification).Error
	return verification, err
}

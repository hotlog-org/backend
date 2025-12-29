package service

import "hotlog.org/models"

type AuthStore interface {
	FindUserByID(id string) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	FindSessionByToken(token string) (models.Session, error)
	FindAccount(providerID, accountID string) (models.Account, error)
	FindVerification(identifier, value string) (models.Verification, error)
}

type AuthService struct {
	store AuthStore
}

func NewAuthService(store AuthStore) *AuthService {
	return &AuthService{store: store}
}

func (s *AuthService) GetUserByID(id string) (models.User, error) {
	return s.store.FindUserByID(id)
}

func (s *AuthService) GetUserByEmail(email string) (models.User, error) {
	return s.store.FindUserByEmail(email)
}

func (s *AuthService) GetSessionByToken(token string) (models.Session, error) {
	return s.store.FindSessionByToken(token)
}

func (s *AuthService) GetAccount(providerID, accountID string) (models.Account, error) {
	return s.store.FindAccount(providerID, accountID)
}

func (s *AuthService) GetVerification(identifier, value string) (models.Verification, error) {
	return s.store.FindVerification(identifier, value)
}

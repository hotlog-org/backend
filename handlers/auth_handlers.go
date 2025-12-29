package handlers

import (
	"errors"
	"net/http"

	"hotlog.org/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/users/:id", h.getUserByID)
	router.GET("/users/by-email", h.getUserByEmail)
	router.GET("/sessions/:token", h.getSessionByToken)
	router.GET("/accounts", h.getAccount)
	router.GET("/verifications", h.getVerification)
}

func (h *AuthHandler) getUserByID(c *gin.Context) {
	user, err := h.service.GetUserByID(c.Param("id"))
	if err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) getUserByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	user, err := h.service.GetUserByEmail(email)
	if err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) getSessionByToken(c *gin.Context) {
	session, err := h.service.GetSessionByToken(c.Param("token"))
	if err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, session)
}

func (h *AuthHandler) getAccount(c *gin.Context) {
	providerID := c.Query("providerId")
	accountID := c.Query("accountId")
	if providerID == "" || accountID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "providerId and accountId are required"})
		return
	}

	account, err := h.service.GetAccount(providerID, accountID)
	if err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *AuthHandler) getVerification(c *gin.Context) {
	identifier := c.Query("identifier")
	value := c.Query("value")
	if identifier == "" || value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "identifier and value are required"})
		return
	}

	verification, err := h.service.GetVerification(identifier, value)
	if err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, verification)
}

func (h *AuthHandler) writeError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

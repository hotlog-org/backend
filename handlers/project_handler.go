package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hotlog.org/service"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func NewProjectHandler(service *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

func (h *ProjectHandler) RegisterRoutes(router gin.IRoutes) {
	router.GET("/projects/:id/events-count", h.getProjectEventsCount)
}

func (h *ProjectHandler) getProjectEventsCount(c *gin.Context) {
	count, err := h.service.GetProjectEventsCount(c.Param("id"))

	if err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, count)
}

func (h *ProjectHandler) writeError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

package repository

import (
	// "hotlog.org/models"

	"gorm.io/gorm"
	"hotlog.org/models"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) GetProjectEventsCount(id string) (int64, error) {
	var model models.Event
	var count int64 = 0

	err := r.db.Where("project_id = ?", id).Model(&model).Count(&count).Error

	return count, err
}

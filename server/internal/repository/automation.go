package repository

import (
	"context"

	"github.com/phzeng0726/go-server-template/internal/domain"

	"gorm.io/gorm"
)

type AutomationsRepo struct {
	db *gorm.DB
}

func NewAutomationsRepo(db *gorm.DB) *AutomationsRepo {
	return &AutomationsRepo{
		db: db,
	}
}

func (r *AutomationsRepo) GetAutomationByParam(ctx context.Context, projectName string, autoType string) ([]domain.Automation, error) {
	var autos []domain.Automation
	db := r.db.WithContext(ctx)

	if err := db.Where("project LIKE ? AND automation_type LIKE ?", projectName, autoType).Find(&autos).Error; err != nil {
		return nil, err
	}

	return autos, nil
}

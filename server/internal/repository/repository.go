package repository

import (
	"context"

	"github.com/phzeng0726/go-server-template/internal/domain"

	"gorm.io/gorm"
)

type Automations interface {
	GetAutomationByParam(ctx context.Context, projectName string, autoType string) ([]domain.Automation, error)
}

type Repositories struct {
	Automations Automations
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Automations: NewAutomationsRepo(db),
	}
}

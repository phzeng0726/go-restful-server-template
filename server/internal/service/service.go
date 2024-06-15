package service

import (
	"context"

	"github.com/phzeng0726/go-server-template/internal/domain"
	"github.com/phzeng0726/go-server-template/internal/repository"
	"github.com/phzeng0726/go-server-template/pkg/auth"
	"github.com/phzeng0726/go-server-template/pkg/logger"
)

type QueryAutomationsInput struct {
	Project string
	Type    string
}

type Automations interface {
	GetIdByParam(ctx context.Context, input QueryAutomationsInput) (domain.Automation, error)
}

type Services struct {
	Automations Automations
}

type Deps struct {
	Repos         *repository.Repositories
	LoggerManager logger.LoggerManager
	TokenManager  auth.TokenManager
}

func NewServices(deps Deps) *Services {
	automationsService := NewAutomationsService(
		deps.Repos.Automations,
		deps.LoggerManager,
	)

	return &Services{
		Automations: automationsService,
	}
}

package service

import (
	"context"
	"errors"

	"github.com/phzeng0726/go-server-template/internal/domain"
	"github.com/phzeng0726/go-server-template/internal/repository"
	"github.com/phzeng0726/go-server-template/pkg/logger"
)

type AutomationsService struct {
	repo          repository.Automations
	loggerManager logger.LoggerManager
}

func NewAutomationsService(
	repo repository.Automations,
	loggerManager logger.LoggerManager,

) *AutomationsService {
	return &AutomationsService{
		repo:          repo,
		loggerManager: loggerManager,
	}
}

func (s *AutomationsService) GetIdByParam(ctx context.Context, input QueryAutomationsInput) (domain.Automation, error) {
	var auto domain.Automation

	autos, err := s.repo.GetAutomationByParam(ctx, input.Project, input.Type)
	if err != nil {
		return auto, err
	}

	if len(autos) == 0 {
		return auto, errors.New("automation not found by the given parameters")
	}

	return autos[0], nil
}

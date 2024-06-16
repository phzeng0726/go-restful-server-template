package service

import (
	"context"

	"github.com/phzeng0726/go-server-template/internal/domain"
	"github.com/phzeng0726/go-server-template/internal/repository"
	"github.com/phzeng0726/go-server-template/pkg/logger"
)

type UsersService struct {
	repo          repository.Users
	loggerManager logger.LoggerManager
}

func NewUsersService(
	repo repository.Users,
	loggerManager logger.LoggerManager,

) *UsersService {
	return &UsersService{
		repo:          repo,
		loggerManager: loggerManager,
	}
}

func (s *UsersService) CreateUser(ctx context.Context, input CreateUserInput) error {
	if err := s.repo.CreateUser(ctx, domain.User{
		Name:  input.Name,
		Email: input.Email,
	}); err != nil {
		return err
	}

	return nil
}

func (s *UsersService) GetUserByEmail(ctx context.Context, input QueryUsersInput) (domain.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

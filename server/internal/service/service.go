package service

import (
	"context"

	"github.com/phzeng0726/go-server-template/internal/domain"
	"github.com/phzeng0726/go-server-template/internal/repository"
	"github.com/phzeng0726/go-server-template/pkg/auth"
	"github.com/phzeng0726/go-server-template/pkg/logger"
)

type CreateUserInput struct {
	Name  string
	Email string
}

type QueryUsersInput struct {
	Email string
}

type Users interface {
	CreateUser(ctx context.Context, input CreateUserInput) error
	GetUserByEmail(ctx context.Context, input QueryUsersInput) (domain.User, error)
}

type Services struct {
	Users Users
}

type Deps struct {
	Repos         *repository.Repositories
	LoggerManager logger.LoggerManager
	TokenManager  auth.TokenManager
}

func NewServices(deps Deps) *Services {
	UsersService := NewUsersService(
		deps.Repos.Users,
		deps.LoggerManager,
	)

	return &Services{
		Users: UsersService,
	}
}

//go:build wireinject

package wire

import (
	"github.com/google/wire"

	"Go/internal/controller"
	"Go/internal/repo"
	"Go/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
	 	service.NewUserService,
	 	controller.NewUserController,
	)

	return new(controller.UserController),nil
}
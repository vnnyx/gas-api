//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	userController "github.com/vnnyx/gas-api/controller/user"
	"github.com/vnnyx/gas-api/infrastructure"
	userRepository "github.com/vnnyx/gas-api/repository/user"
	userService "github.com/vnnyx/gas-api/service/user"
)

func InitializeUserController(configName string) userController.UserController {
	wire.Build(
		infrastructure.NewConfig,
		infrastructure.NewMySQLDatabase,
		userRepository.NewUserRepository,
		userService.NewUserService,
		userController.NewUserController,
	)
	return nil
}

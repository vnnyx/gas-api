package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vnnyx/gas-api/exception"
	"github.com/vnnyx/gas-api/infrastructure"
	"github.com/vnnyx/gas-api/injector/wire"
	"github.com/vnnyx/gas-api/migration"
	"github.com/vnnyx/gas-api/model"
)

func main() {
	configuration := infrastructure.NewConfig(".env")
	database := infrastructure.NewMySQLDatabase(configuration)
	migration.Migrate(database, model.User{}, model.Fashion{}, model.Kelontong{})

	userController := wire.InitializeUserController(".env")

	app := echo.New()
	app.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{DisablePrintStack: true}))
	app.Use(middleware.CORS())
	app.HTTPErrorHandler = exception.ErrorHandler
	userController.Route(app)
	err := app.Start(fmt.Sprintf(":%v", configuration.AppPort))
	exception.PanicIfNeeded(err)
}

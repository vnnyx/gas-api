package user

import (
	"github.com/labstack/echo/v4"
)

type UserController interface {
	Route(e *echo.Echo)
	CreateUser(c echo.Context) error
	GetUserById(c echo.Context) error
	GetAllUSer(c echo.Context) error
	UpdateUserProfile(c echo.Context) error
	RemoveUser(c echo.Context) error
}

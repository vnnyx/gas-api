package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vnnyx/gas-api/exception"
	"github.com/vnnyx/gas-api/model"
	"github.com/vnnyx/gas-api/service/user"
)

type UserControllerImpl struct {
	user.UserService
}

func NewUserController(userService user.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) Route(e *echo.Echo) {
	api := e.Group("/gas-api/user")
	api.POST("", controller.CreateUser)
	api.GET("/:id", controller.GetUserById)
	api.GET("", controller.GetAllUSer)
	api.PUT("/:id", controller.UpdateUserProfile)
	api.DELETE("/:id", controller.RemoveUser)
}

func (controller *UserControllerImpl) CreateUser(c echo.Context) error {
	var request model.UserCreateRequest
	err := c.Bind(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.UserService.CreateUser(c.Request().Context(), request)
	exception.PanicIfNeeded(err)

	return c.JSON(http.StatusCreated, model.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *UserControllerImpl) GetUserById(c echo.Context) error {
	userID := c.Param("id")

	response, err := controller.UserService.GetUserByID(c.Request().Context(), userID)
	exception.PanicIfNeeded(err)

	return c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserControllerImpl) GetAllUSer(c echo.Context) error {
	response, err := controller.UserService.GetAllUser(c.Request().Context())
	exception.PanicIfNeeded(err)

	return c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserControllerImpl) UpdateUserProfile(c echo.Context) error {
	var request model.UserUpdateRequest
	err := c.Bind(&request)
	exception.PanicIfNeeded(err)

	request.ID = c.Param("id")
	response, err := controller.UserService.UpdateUserProfile(c.Request().Context(), request)

	return c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserControllerImpl) RemoveUser(c echo.Context) error {
	userID := c.Param("id")

	err := controller.UserService.RemoveUser(c.Request().Context(), userID)
	exception.PanicIfNeeded(err)

	return c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	})
}

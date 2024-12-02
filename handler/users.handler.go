package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nrmnqdds/vtech-qms-be/db/store"
)

func (app *App) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	logger := app.Logger.GetLogger()

	logger.Debugf("Creating user with request: %v", c.Request().Body)
	e := new(store.CreateUserParams)
	err := c.Bind(e)
	if err != nil {
		logger.Errorf("Failed to bind request: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind request")
	}

	res, err := app.Queries.CreateUser(ctx, *e)
	if err != nil {
		logger.Errorf("Failed to create user: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, "Failed to create user")
	}

	return c.JSON(http.StatusOK, res)
}

func (app *App) GetAllUsers(c echo.Context) error {
	ctx := c.Request().Context()

	logger := app.Logger.GetLogger()

	logger.Debug("Getting all users")
	res, err := app.Queries.GetAllUsers(ctx)
	if err != nil {
		logger.Errorf("Failed to get all users: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, "Failed to get all users")
	}

	return c.JSON(http.StatusOK, res)
}

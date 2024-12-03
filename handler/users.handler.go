package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nrednav/cuid2"
	"github.com/nrmnqdds/vtech-qms-be/db/store"
)

// UserHandler is an interface for user handler
type UserHandler interface {
	CreateUser(c echo.Context) error
	GetAllUsers(c echo.Context) error
}

// UserHandlerImpl implements UserHandler interface
type UserHandlerImpl struct {
	app *App
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(app *App) UserHandler {
	return &UserHandlerImpl{
		app: app,
	}
}

// CreateUser creates a new user
func (h *UserHandlerImpl) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	logger := h.app.Logger.GetLogger()

	var err error

	generateCUID, err := cuid2.Init(
		cuid2.WithLength(16),
	)
	if err != nil {
		logger.Errorf("Failed to generate cuid: %v", err)
	}

	logger.Debugf("Creating user with request: %v", c.Request().Body)
	e := new(store.CreateUserParams)
	e.ID = generateCUID()
	err = c.Bind(e)
	if err != nil {
		logger.Errorf("Failed to bind request: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind request")
	}

	res, err := h.app.Queries.CreateUser(ctx, *e)
	if err != nil {
		logger.Errorf("Failed to create user: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, "Failed to create user")
	}

	return c.JSON(http.StatusOK, res)
}

// GetAllUsers gets all users
func (h *UserHandlerImpl) GetAllUsers(c echo.Context) error {
	ctx := c.Request().Context()

	logger := h.app.Logger.GetLogger()

	logger.Debug("Getting all users")
	res, err := h.app.Queries.GetAllUsers(ctx)
	if err != nil {
		logger.Errorf("Failed to get all users: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, "Failed to get all users")
	}

	return c.JSON(http.StatusOK, res)
}

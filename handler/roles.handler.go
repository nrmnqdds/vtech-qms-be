package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RoleHandler is an interface for user handler
type RoleHandler interface {
	SeedRoles(c echo.Context) error
}

// RoleHandlerImpl implements RoleHandler interface
type RoleHandlerImpl struct {
	app *App
}

// NewRoleHandler creates a new RoleHandler instance
func NewRoleHandler(app *App) RoleHandler {
	return &RoleHandlerImpl{
		app: app,
	}
}

func (h *RoleHandlerImpl) SeedRoles(c echo.Context) error {
	ctx := c.Request().Context()

	if err := h.app.Queries.SeedRoles(ctx); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, "Roles seeded!")
}

package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *App) SeedRoles(c echo.Context) error {
	ctx := c.Request().Context()

	if err := app.Queries.SeedRoles(ctx); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, "Roles seeded!")
}

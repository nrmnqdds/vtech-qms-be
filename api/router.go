package api

import (
	"github.com/nrmnqdds/vtech-qms-be/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserRouter(g *echo.Group, app *handler.App) {
	g.GET("", app.GetAllUsers)
	g.POST("/create", app.CreateUser)
}

func RoleRouter(g *echo.Group, app *handler.App) {
	g.POST("/seed", app.SeedRoles)
}

//	func RoutineRouter(g *echo.Group, app *handler.App) {
//		g.GET("", app.GetRoutineList)
//		g.POST("", app.CreateRoutine)
//		g.GET("/:id", app.GetRoutine)
//		g.PUT("/:id", app.UpdateRoutine)
//		g.DELETE("/:id", app.DeleteRoutine)
//	}

func Router(e *echo.Echo, app *handler.App) {
	e.GET("/ping", handler.PingHandler)

	UserRouter(e.Group("/api/users", middleware.KeyAuth(CheckAuthHeader)), app)
	RoleRouter(e.Group("/api/roles", middleware.KeyAuth(CheckAuthHeader)), app)
}

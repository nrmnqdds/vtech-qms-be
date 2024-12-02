package main

import (
	"context"

	"github.com/nrmnqdds/vtech-qms-be/api"
	"github.com/nrmnqdds/vtech-qms-be/config"
	"github.com/nrmnqdds/vtech-qms-be/config/logger"
	"github.com/nrmnqdds/vtech-qms-be/db"
	"github.com/nrmnqdds/vtech-qms-be/handler"

	"github.com/brpaz/echozap"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(); err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	connPool := db.GetConnPool(ctx)
	defer connPool.Close(ctx)

	logger := logger.New()

	e.Use(echozap.ZapLogger(logger.Logger))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper:          middleware.DefaultSkipper,
		Format:           "${method} ${status} ${uri} ${remote_ip} ${user_agent}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           e.Logger.Output(),
	}))

	app := handler.NewApp(connPool)
	api.Router(e, app)

	e.Logger.Fatal(e.Start(config.GetServerString()))
}

package prepare

import (
	"github.com/DanielTitkov/correlateme-server/internal/api/handler"
	"github.com/DanielTitkov/correlateme-server/internal/app"
	"github.com/DanielTitkov/correlateme-server/internal/configs"
	"github.com/DanielTitkov/correlateme-server/internal/logger"
	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewServer(cfg configs.Config, logger *logger.Logger, app *app.App) *echo.Echo {
	e := echo.New()
	// Enable metrics middleware
	e.Use(echoPrometheus.MetricsMiddleware())
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	if cfg.Env != "dev" {
		e.Use(middleware.Recover())
	}
	handler.NewHandler(e, cfg, logger, app)
	return e
}

package handler

import (
	"github.com/DanielTitkov/correlateme-server/internal/app"
	"github.com/DanielTitkov/correlateme-server/internal/configs"
	"github.com/DanielTitkov/correlateme-server/internal/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Handler struct {
	cfg    configs.Config
	logger *logger.Logger
	app    *app.App
}

func NewHandler(
	e *echo.Echo,
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) *Handler {
	h := &Handler{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
	h.link(e)
	return h
}

func (h *Handler) link(e *echo.Echo) {

	v1 := e.Group("/api/v1")
	v1.POST("/getToken", h.GetTokenHandler)
	v1.POST("/createUser", h.CreateUserHandler)
	// Restricted group only with valid JWT
	v1Restricted := v1.Group("/private")
	v1Restricted.Use(middleware.JWT([]byte(h.cfg.Auth.Secret)))
	v1Restricted.POST("/getUser", h.GetUserHandler)

	// TODO
	// v1Restricted.POST("/", h.CreateIndicator)
	// v1Restricted.POST("/", h.GetIndicators)
	// v1Restricted.POST("/", h.CreateObservation)
	// v1Restricted.POST("/", h.GetDataset)
	// v1Restricted.POST("/", h.GetCorrelation) // ?
	// v1Restricted.POST("/", h.GetCorrelations) // ?
	// v1Restricted.POST("/", h.)
	// v1Restricted.POST("/", h.)
}

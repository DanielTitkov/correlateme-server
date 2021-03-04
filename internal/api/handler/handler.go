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
	// restricted group only with valid JWT
	v1Restricted := v1.Group("/private")
	v1Restricted.Use(middleware.JWT([]byte(h.cfg.Auth.Secret)))
	v1Restricted.POST("/getUser", h.GetUserHandler)
	// indicator
	v1Restricted.POST("/createIndicator", h.CreateIndicator)
	v1Restricted.POST("/getIndicators", h.GetIndicators)
	// observation
	v1Restricted.POST("/createOrUpdateObservation", h.CreateOrUpdateObservation)
	// correlation
	v1Restricted.POST("/getCorrelationMatrix", h.GetCorrelationMatrix)

	// TODO
	// v1Restricted.POST("/", h.GetDataset)
	// v1Restricted.POST("/", h.GetCorrelation) // ?
	// v1Restricted.POST("/", h.GetCorrelations) // ?
	// v1Restricted.POST("/", h.)
	// v1Restricted.POST("/", h.)

	// dev endpoints for testing and debug
	// not accessible on not-dev instances
	if h.cfg.Env == "dev" {
		v1.POST("/findUserCorrelations", h.FindUserCorrelations)
	}
}

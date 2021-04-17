package handler

import (
	"net/http"

	"github.com/DanielTitkov/correlateme-server/internal/api/model"
	"github.com/labstack/echo"
)

func (h *Handler) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "service is running",
	})
}

func (h *Handler) StatsHandler(c echo.Context) error {
	stats, err := h.app.ServiceStats()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to get service stats",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.ServiceStatsResponse{
		Users:        stats.Users,
		Scales:       stats.Scales,
		Indicators:   stats.Indicators,
		Datasets:     stats.Datasets,
		Observations: stats.Observations,
		Correlations: stats.Correlations,
	})
}

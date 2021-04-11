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

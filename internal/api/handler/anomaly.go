package handler

import (
	"net/http"

	"github.com/DanielTitkov/correlateme-server/internal/api/model"
	"github.com/labstack/echo"
)

func (h *Handler) ListAnomaliesHandler(c echo.Context) error {
	request := new(model.ListAnomaliesRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	anomalies, err := h.app.ListAnomalies(&request.Filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to filter anomalies",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.ListAnomaliesResponse{
		Anomalies: anomalies,
	})
}

func (h *Handler) SetAnomalyStatusHandler(c echo.Context) error {
	request := new(model.SetAnomalyStatusRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "",
	})
}

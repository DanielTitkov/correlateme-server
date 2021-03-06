package handler

import (
	"net/http"

	"github.com/DanielTitkov/correlateme-server/internal/api/model"
	"github.com/DanielTitkov/correlateme-server/internal/api/util"
	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/labstack/echo"
)

func (h *Handler) CreateOrUpdateObservation(c echo.Context) error {
	request := new(model.CreateOrUpdateObservationRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	userID, err := util.UserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "token is invalid",
			Error:   err.Error(),
		})
	}

	err = h.app.CreateOrUpdateObservation(domain.CreateOrUpdateObservationArgs{
		UserID:      userID,
		IndicatorID: request.IndicatorID,
		Value:       request.Value,
		Date:        request.Date,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to create/update observation",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "observation created/updated",
	})
}

func (h *Handler) UpdateAggregations(c echo.Context) error {
	request := new(model.UpdateAggregationsRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	if request.DatasetID == 0 {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed",
			Error:   "datasetID is not provided",
		})
	}

	err := h.app.UpdateAggregations(domain.UpdateAggregationsArgs{
		DatasetID: request.DatasetID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "done",
	})
}

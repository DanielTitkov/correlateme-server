package handler

import (
	"net/http"

	"github.com/DanielTitkov/correlateme-server/internal/api/model"
	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/labstack/echo"
)

// FindUserCorrelations method is only for debug on dev
func (h *Handler) FindUserCorrelations(c echo.Context) error {
	request := new(model.FindUserCorrelationsRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	if request.UserID == 0 {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed",
			Error:   "userID is not provided",
		})
	}

	err := h.app.FindCorrelations(domain.FindCorrelationsArgs{
		UserID:     request.UserID,
		WithShared: request.WithShared,
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

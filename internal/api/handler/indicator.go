package handler

import (
	"net/http"

	"github.com/DanielTitkov/correlateme-server/internal/api/model"
	"github.com/DanielTitkov/correlateme-server/internal/api/util"
	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/labstack/echo"
)

func (h *Handler) CreateIndicator(c echo.Context) error {
	request := new(model.CreateIndicatorRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	// TODO: add validation
	// if err := c.Validate(request); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	username, err := util.UsernameFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "token is invalid",
			Error:   err.Error(), // TODO: bind with echo logging
		})
	}

	err = h.app.CreateIndicator(domain.CreateIndicatorArgs{
		Username:    username,
		ScaleType:   request.ScaleType,
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to create indicator",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "indicator created",
	})
}

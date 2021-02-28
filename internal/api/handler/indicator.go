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

func (h *Handler) GetIndicators(c echo.Context) error {
	request := new(model.GetIndicatorsRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	username, err := util.UsernameFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "token is invalid",
			Error:   err.Error(), // TODO: bind with echo logging
		})
	}

	filter := request.Filter
	indicators, err := h.app.FilterIndicators(domain.FilterIndicatorsArgs{
		AuthorUsername: &username,
		ID:             filter.ID,
		Code:           filter.Code,
		Title:          filter.Title,
		Active:         filter.Active,
		BuiltIn:        filter.BuiltIn,
		ScaleType:      filter.ScaleType,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to get indicators",
			Error:   err.Error(),
		})
	}

	response := model.GetIndicatorsResponse{}
	for _, i := range indicators {
		var authorID int
		if i.Author != nil {
			authorID = i.Author.ID
		}
		response = append(response, model.GetIndicatorsResponseItem{
			ID:          i.ID,
			Code:        i.Code,
			Title:       i.Title,
			Description: i.Description,
			Active:      i.Active,
			BuiltIn:     i.BuiltIn,
			External:    i.External,
			ScaleID:     i.Scale.ID,
			AuthorID:    authorID,
			CreateTime:  i.CreateTime,
			UpdateTime:  i.UpdateTime,
		})
	}

	return c.JSON(http.StatusOK, response)
}
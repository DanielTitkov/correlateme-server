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
		Username:     username,
		ScaleType:    request.ScaleType,
		Title:        request.Title,
		Description:  request.Description,
		ValueMapping: request.ValueMapping,
		ValueParams:  request.ValueParams,
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

func (h *Handler) UpdateIndicator(c echo.Context) error {
	request := new(model.UpdateIndicatorRequest)
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

	err = h.app.UpdateIndicator(domain.UpdateIndicatorArgs{
		UserID:         userID,
		ID:             request.ID,
		Active:         request.Active,
		Title:          request.Title,
		Description:    request.Description,
		ValueMapping:   request.ValueMapping,
		ValueParams:    request.ValueParams,
		UpdateBuiltins: false,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to update indicator",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "indicator updated",
	})
}

func (h *Handler) GetIndicators(c echo.Context) error {
	request := new(model.GetIndicatorsRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	userID, err := util.UserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "token is invalid",
			Error:   err.Error(), // TODO: bind with echo logging
		})
	}

	filter := request.Filter
	indicators, err := h.app.GetIndicators(domain.GetIndicatorsArgs{
		UserID:           userID,
		WithDataset:      request.WithDataset,
		ObservationLimit: request.ObservationsLimit,
		Granularity:      request.Granularity,
		Filter: domain.GetIndicatorsArgsFilter{
			AuthorID:  &userID,
			ID:        filter.ID,
			Code:      filter.Code,
			Title:     filter.Title,
			Active:    filter.Active,
			BuiltIn:   filter.BuiltIn,
			ScaleType: filter.ScaleType,
		},
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
		var dataset *model.Dataset
		if i.UserDataset != nil {
			var observations []model.Observation
			if i.UserDataset.Observations != nil && len(i.UserDataset.Observations) > 0 {
				for _, obs := range i.UserDataset.Observations {
					observations = append(observations, model.Observation{
						ID:    obs.ID,
						Value: obs.Value,
						Date:  obs.Date,
					})
				}
			}
			dataset = &model.Dataset{
				ID:           i.UserDataset.ID,
				Observations: observations,
				Granularity:  request.Granularity,
			}
			if i.UserDataset.Params != nil {
				dataset.Style = i.UserDataset.Params.Style
				dataset.Aggregation = i.UserDataset.Params.Aggregation
			}
		}

		response = append(response, model.GetIndicatorsResponseItem{
			Indicator: model.Indicator{
				ID:           i.ID,
				Code:         i.Code,
				Title:        i.Title,
				Description:  i.Description,
				Active:       i.Active,
				BuiltIn:      i.BuiltIn,
				External:     i.External,
				ScaleID:      i.Scale.ID,
				AuthorID:     authorID,
				CreateTime:   i.CreateTime,
				UpdateTime:   i.UpdateTime,
				Dataset:      dataset,
				ValueMapping: i.ValueMapping,
				ValueParams:  i.ValueParams,
			},
		})
	}

	return c.JSON(http.StatusOK, response)
}

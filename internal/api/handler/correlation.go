package handler

import (
	"net/http"

	"github.com/DanielTitkov/correlateme-server/internal/api/model"
	"github.com/DanielTitkov/correlateme-server/internal/api/util"
	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/labstack/echo"
)

func (h *Handler) GetCorrelationMatrix(c echo.Context) error {
	request := new(model.GetCorrelationMatrixRequest)
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

	granularity := domain.GranularityDay
	if request.Granularity != "" {
		granularity = request.Granularity
	}

	matrix, err := h.app.GetCorrelationMatrix(domain.GetCorrelationMatrixArgs{
		UserID:      userID,
		WithShared:  request.WithShared,
		Granularity: granularity,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to get correlation matrix",
			Error:   err.Error(),
		})
	}

	var header []model.GetCorrelationMatrixResponseHeaderItem
	var body [][]model.GetCorrelationMatrixResponseBodyItem

	for _, hi := range matrix.Header {
		header = append(header, model.GetCorrelationMatrixResponseHeaderItem{
			IndicatorID:    hi.IndicatorID,
			DatasetID:      hi.DatasetID,
			IndicatorTitle: hi.IndicatorTitle,
			DatasetShared:  hi.DatasetShared,
		})
	}

	for _, br := range matrix.Body {
		var bodyRow []model.GetCorrelationMatrixResponseBodyItem
		for _, bi := range br {
			bodyRow = append(bodyRow, model.GetCorrelationMatrixResponseBodyItem{
				CorrelationID: bi.CorrelationID,
				Coef:          bi.Coef,
				P:             bi.P,
				R2:            bi.R2,
				Type:          bi.Type,
				UpdateTime:    bi.UpdateTime,
			})
		}
		body = append(body, bodyRow)
	}

	return c.JSON(http.StatusOK, model.GetCorrelationMatrixResponse{
		Granularity: granularity,
		Header:      header,
		Body:        body,
	})
}

func (h *Handler) GetCorrelation(c echo.Context) error {
	request := new(model.GetCorrelationRequest)
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

	// TODO: check if user exists?

	corr, err := h.app.GetCorrelation(domain.GetCorrelationArgs{
		ID:               request.ID,
		UserID:           userID,
		WithDatasets:     true,
		WithObservations: true,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to get correlation",
			Error:   err.Error(),
		})
	}

	response := model.GetCorrelationResponse{
		Correlation: model.Correlation{
			ID:          corr.ID,
			Coef:        corr.Coef,
			P:           corr.P,
			R2:          corr.R2,
			Type:        corr.Type,
			Granularity: corr.Granularity,
			UpdateTime:  corr.UpdateTime,
			Left:        domainToApiDataset(corr.Left, corr.Granularity),
			Right:       domainToApiDataset(corr.Right, corr.Granularity),
		},
	}

	return c.JSON(http.StatusOK, response)
}

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

	err := h.app.UpdateCorrelations(domain.UpdateCorrelationsArgs{
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

func domainToApiDataset(ds *domain.Dataset, granularity string) *model.Dataset {
	if ds == nil {
		return nil
	}

	var observations []model.Observation
	if ds.Observations != nil {
		for _, obs := range ds.Observations {
			observations = append(observations, model.Observation{
				ID:    obs.ID,
				Value: obs.Value,
				Date:  obs.Date,
			})
		}
	}

	return &model.Dataset{
		ID:           ds.ID,
		Source:       ds.Source,
		Shared:       ds.Shared,
		Granularity:  granularity,
		Observations: observations,
	}
}

package app

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/dgryski/go-onlinestats"
	combinations "github.com/mxschmitt/golang-combinations"
)

func (a *App) GetCorrelation(args domain.GetCorrelationArgs) (*domain.Correlation, error) {
	corr, err := a.repo.GetCorrelation(args)
	if err != nil {
		return nil, err
	}

	return corr, nil
}

func (a *App) GetCorrelationMatrix(args domain.GetCorrelationMatrixArgs) (*domain.CorrelationMatrix, error) {
	datasets, err := a.repo.GetDatasets(domain.GetDatasetsArgs{
		UserID:           args.UserID,
		WithIndicator:    true,
		ObservationLimit: 0,
		Filter: domain.GetDatasetsArgsFilter{
			WithShared: args.WithShared,
		},
	})
	if err != nil {
		return nil, err
	}

	correlations, err := a.repo.GetUserCorrelations(args.UserID)
	if err != nil {
		return nil, err
	}

	// map correlations
	corrMap := mapCorrelation(correlations)

	// make matrix
	var matrix [][]domain.CorrelationMatrixBodyItem
	for _, iDataset := range datasets {
		var matrixRow []domain.CorrelationMatrixBodyItem
		for _, jDataset := range datasets {
			if iDataset.ID == jDataset.ID {
				matrixRow = append(matrixRow, makeSelfCorrelationMatrixItem())
				continue
			}
			corr, ok := corrMap[pairOfIDsToString(iDataset.ID, jDataset.ID)]
			if !ok {
				matrixRow = append(matrixRow, makeZeroCorrelationMatrixItem())
				continue
			}
			matrixRow = append(matrixRow, makeCorrelationMatrixItem(corr))
		}
		matrix = append(matrix, matrixRow)
	}

	var header []domain.CorrelationMatrixHeaderItem
	for _, dataset := range datasets {
		header = append(header, domain.CorrelationMatrixHeaderItem{
			IndicatorID:    dataset.Indicator.ID,
			DatasetID:      dataset.ID,
			IndicatorTitle: dataset.Indicator.Title,
			DatasetShared:  dataset.Shared,
		})
	}

	return &domain.CorrelationMatrix{
		Header: header,
		Body:   matrix,
	}, nil
}

func (a *App) UpdateCorrelations(args domain.UpdateCorrelationsArgs) error {
	// day // FIXME
	datasets, err := a.repo.GetUserDatasets(args.UserID, args.WithShared, int(a.cfg.App.MaxCorrelationObservations), "day")
	if err != nil {
		return err
	}

	// map datasets
	datasetMap := mapDatasets(datasets) // TODO: create combination method for ints to remove conversion

	// get dataset ids
	var datasetIDs []string
	for k := range datasetMap {
		datasetIDs = append(datasetIDs, k)
	}

	// make combinations
	subsets := combinations.Combinations(datasetIDs, 2)

	// make data slices
	for _, subset := range subsets {
		left := datasetMap[subset[0]]
		right := datasetMap[subset[1]]
		leftObs := mapObservations(left.Observations)
		rightObs := mapObservations(right.Observations)

		var leftData, rightData []float64
		for date, leftOb := range leftObs {
			rightOb, ok := rightObs[date]
			if !ok {
				continue
			}
			leftData = append(leftData, leftOb.Value)
			rightData = append(rightData, rightOb.Value)
		}
		if len(leftData) == 0 { // TODO: check minimal observation count
			continue
		}

		// find coef and p
		coef, p := onlinestats.Spearman(leftData, rightData)
		if math.IsNaN(coef) {
			continue
		}
		correlation := domain.Correlation{
			Left:  left,
			Right: right,
			Coef:  coef,
			P:     p,
			Type:  domain.SpearmanCorrelationType,
			R2:    math.Pow(coef, 2),
		}

		// save correlations to db
		_, err = a.repo.CreateOrUpdateCorrelation(&correlation)
		if err != nil { // TODO: maybe save error, not exit right away
			return err
		}
	}

	return nil
}

func mapObservations(obs []*domain.Observation) map[time.Time]*domain.Observation {
	res := make(map[time.Time]*domain.Observation)
	for _, ob := range obs {
		res[*ob.Date] = ob
	}
	return res
}

func mapDatasets(datasets []*domain.Dataset) map[string]*domain.Dataset {
	datasetMap := make(map[string]*domain.Dataset)
	for _, dataset := range datasets {
		if len(dataset.Observations) == 0 {
			continue
		}
		datasetMap[fmt.Sprintf("%d", dataset.ID)] = dataset
	}
	return datasetMap
}

func mapCorrelation(corrs []*domain.Correlation) map[string]*domain.Correlation {
	corrMap := make(map[string]*domain.Correlation)
	for _, corr := range corrs {
		// this way map size is doubled but it provides results for every order
		corrMap[pairOfIDsToString(corr.Left.ID, corr.Right.ID)] = corr
		corrMap[pairOfIDsToString(corr.Right.ID, corr.Left.ID)] = corr
	}
	return corrMap
}

func pairOfIDsToString(left, right int) string {
	return strings.Join([]string{
		strconv.Itoa(left),
		strconv.Itoa(right),
	}, "_") // separator is for readability
}

func makeSelfCorrelationMatrixItem() domain.CorrelationMatrixBodyItem {
	return domain.CorrelationMatrixBodyItem{
		CorrelationID: 0,
		Coef:          1,
		P:             0,
		R2:            0,
		Type:          "-",
		UpdateTime:    time.Now(),
	}
}

func makeZeroCorrelationMatrixItem() domain.CorrelationMatrixBodyItem {
	return domain.CorrelationMatrixBodyItem{
		CorrelationID: 0,
		Coef:          0,
		P:             0,
		R2:            0,
		Type:          "-",
		UpdateTime:    time.Now(),
	}
}

func makeCorrelationMatrixItem(corr *domain.Correlation) domain.CorrelationMatrixBodyItem {
	return domain.CorrelationMatrixBodyItem{
		CorrelationID: corr.ID,
		Coef:          corr.Coef,
		P:             corr.P,
		R2:            corr.R2,
		Type:          corr.Type,
		UpdateTime:    corr.UpdateTime,
	}
}
